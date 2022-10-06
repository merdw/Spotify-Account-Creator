package main

import (
	"fmt"
	browser "github.com/eddycjy/fake-useragent"
	"github.com/mazen160/go-random"
	"github.com/valyala/fasthttp"
	"os"
	"regexp"
	"strings"
)

var (
	dggun, dgay, email, password, gender, nick, logintoken, username, mert string
	dgyil, manyacc                                                         int
)

func main() {

	fmt.Print("How many account will be created?: ")
	fmt.Scanln(&manyacc)

	for i := 0; i < manyacc; i++ {
		runu()
		fmt.Printf("Created %d. account\n", i+1)
	}
	fmt.Printf("Finished Created %d Accounts", manyacc)
}

func runu() {

	infocek()
	kayit()

}
func infocek() {
	client := &fasthttp.Client{}
	client.ReadBufferSize = 8192
	req := &fasthttp.Request{}
	resp := &fasthttp.Response{}
	req.SetRequestURI("https://randomuser.me/api/1.3/?nat=tr")
	req.Header.SetMethod("GET")
	err := client.Do(req, resp)
	if err != nil {
		fmt.Println(err)
		return
	}
	buffer := string(resp.Body())

	deneme := buffer
	re := regexp.MustCompile(`,"first":"(.*)","last`)
	match := re.FindStringSubmatch(deneme)
	isim := match[1]

	rer := regexp.MustCompile(`"last":"(.*)"},"location`)
	matchh := rer.FindStringSubmatch(deneme)
	soyad := matchh[1]

	rere := regexp.MustCompile(`"email":"(.*)","lo`)
	matchhh := rere.FindStringSubmatch(deneme)
	emaillo := matchhh[1]

	rerer := regexp.MustCompile(`(.*)@`)
	matchhhh := rerer.FindStringSubmatch(emaillo)
	emaill := matchhhh[1]

	charset := "abcdefghjklmnoprstuvyz0123456789"
	length := 4
	data, err := random.Random(length, charset, true)
	if err != nil {
		fmt.Println(err)
	}

	password = emaill + data

	email = password + "@gmail.com"

	dgyil, _ = random.IntRange(1990, 2003)

	s := []string{"01", "02", "03", "04", "05", "06", "07", "08", "09", "10", "11", "12"}
	dgay, _ = random.Choice(s)
	dgay = dgay

	f := []string{"11", "12", "13", "14", "15", "16", "17", "18", "19", "20", "21", "22", "23", "24"}
	dggun, _ = random.Choice(f)
	dggun = dggun
	r := []string{"1", "2"}
	gender, _ = random.Choice(r)
	gender = gender
	nick = isim + " " + soyad
}

func kayit() {

	randuser := browser.MacOSX()

	client := &fasthttp.Client{}
	client.ReadBufferSize = 8192
	req := &fasthttp.Request{}
	resp := &fasthttp.Response{}
	req.SetRequestURI("https://spclient.wg.spotify.com/signup/public/v2/account/create")
	req.Header.SetMethod("POST")

	req.Header.Add("User-Agent", randuser)
	req.Header.Add("origin", "https://www.spotify.com")
	req.Header.Add("Pragma", "no-cache")
	req.Header.Add("Accept", "*/*")
	req.Header.Add("accept-language", "tr-TR,tr;q=0.9")
	req.Header.Add("Content-Type", "application/json")

	req.Header.Add("referer", "https://www.spotify.com")
	req.SetBody([]byte(fmt.Sprintf("{\"account_details\":{\"birthdate\":\"%d-%s-%s\",\"consent_flags\":{\"eula_agreed\":true,\"send_email\":true,\"third_party_email\":false},\"display_name\":\"%s\",\"email_and_password_identifier\":{\"email\":\"%s\",\"password\":\"%s\"},\"gender\":%s},\"callback_uri\":\"https://www.spotify.com/signup/challenge?locale=tr\",\"client_info\":{\"api_key\":\"a1e486e2729f46d6bb368d6b2bcda326\",\"app_version\":\"v2\",\"capabilities\":[1],\"installation_id\":\"dc2dbbf4-029e-4983-8d69-a30abe65b8f4\",\"platform\":\"www\"},\"tracking\":{\"creation_flow\":\"\",\"creation_point\":\"https://www.spotify.com/tr/\",\"referrer\":\"\"}}", dgyil, dgay, dggun, nick, email, password, gender)))
	err := client.Do(req, resp)
	if err != nil {
		fmt.Println(err)
	}

	buffer := string(resp.Body())
	if strings.Contains(buffer, "{\"success\":") {
		tokka := regexp.MustCompile(`{"username":"(.*)","login`)
		match := tokka.FindStringSubmatch(buffer)
		username = match[1]
		fmt.Println("username:" + username)

		tokkaa := regexp.MustCompile(`"login_token":"(.*)"}}`)
		matchh := tokkaa.FindStringSubmatch(buffer)
		logintoken = matchh[1]
		fmt.Println("token:" + logintoken)

		emailpwyazdir()
		tokenyazdir()
		return
	} else {
		fmt.Println("ERROR")
	}
}

func emailpwyazdir() {
	file, err := os.OpenFile("createdspotify.txt", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	file.WriteString(fmt.Sprintf("%s:%s\n", email, password))
	if err != nil {
		fmt.Println(err)
	} else {
	}
}
func tokenyazdir() {
	file, err := os.OpenFile("createdspotifytokens.txt", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	file.WriteString(fmt.Sprintf("%s:%s\n", logintoken, username))
	if err != nil {
		fmt.Println(err)
	} else {
	}
}
