package bookmyvaccine

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"github.com/patilsuraj767/bookmyvaccine/bookmyvaccine/model"
	"github.com/spf13/viper"
)

func CheckAvailability() {
	var resData model.Data
	appointmentLink := "https://cdn-api.co-vin.in/api/v2/appointment/sessions/public/calendarByPin?pincode=" + viper.GetString("PIN_CODE") + "&date=" + viper.GetString("DATE")

	f, _ := os.Open("bookmyvaccine/files/dubstep.mp3")
	streamer, format, _ := mp3.Decode(f) // yum install alsa-lib-devel
	defer streamer.Close()

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	req, _ := http.NewRequest("GET", appointmentLink, nil)
	client := &http.Client{}
	req.Header.Add("accept", "application/json")
	req.Header.Add("Accept-Language", "en_US")
	req.Header.Add("User-Agent", "Mozilla/5.0 (X11; Fedora; Linux x86_64; rv:70.0) Gecko/20100101 Firefox/70.0")
	for {
		fmt.Printf("Checking vaccine availability age group greater than %d at location %v \n", viper.GetInt("AGE_GROUP"), viper.GetInt("PIN_CODE"))
		resp, _ := client.Do(req)
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		json.Unmarshal(body, &resData)
		//fmt.Println(resData.Centers)
		for _, c := range resData.Centers {
			for _, s := range c.Sessions {
				if s.Min_age_limit == viper.GetInt("AGE_GROUP") && s.Available_capacity >= viper.GetInt("AVAILABLE_CAPACITY") {
					vaccineAvailableCenter, _ := json.Marshal(c)
					fmt.Println(string(vaccineAvailableCenter))
					speaker.Play(streamer)
					select {}
				}
			}
		}
		time.Sleep(30 * time.Second)
	}
}
