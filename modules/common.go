package modules

import (
	"strconv"

	"github.com/adlio/trello"
)

func init() {

}

func GetRealTimeOfDone(name string) int {
	l := len(name)
	time := ""
	for i := l - 1; i > 0; i-- {
		if string(name[i]) == "]" {
			i--
			for ; string(name[i]) != "["; i-- {
				time = string(name[i]) + string(time)
			}
			break
		}
	}
	ret, _ := strconv.Atoi(time)
	return ret
}

func GetTimeGuessForDone(name string) int {
	l := len(name)
	time := ""
	for i := l - 1; i > 0; i-- {
		if string(name[i]) == ")" {
			i--
			for ; string(name[i]) != "("; i-- {
				time = string(name[i]) + string(time)
			}
			break
		}
	}
	ret, _ := strconv.Atoi(time)
	return ret
}

func ProvideMemberAndCLient(key, tk, id string) (*trello.Client, *trello.Member, error) {
	client := trello.NewClient(key, tk)
	member, err := client.GetMember(id, trello.Defaults())
	if err != nil {
		return client, nil, err
	}
	return client, member, nil
}
