package pkg

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

var champs []Champion

type Champion struct {
	ID    string   `json:"id"`
	Key   string   `json:"key"`
	Name  string   `json:"name"`
	Title string   `json:"title"`
	Tags  []string `json:"tags"`
	Stats struct {
		HP                   float64 `json:"hp"`
		HpPerLevel           float64 `json:"hpperlevel"`
		MP                   float64 `json:"mp"`
		MpPerLevel           float64 `json:"mpperlevel"`
		MoveSpeed            float64 `json:"movespeed"`
		Armor                float64 `json:"armor"`
		ArmorPerLevel        float64 `json:"armorperlevel"`
		SpellBlock           float64 `json:"spellblock"`
		SpellBlockPerLevel   float64 `json:"spellblockperlevel"`
		AttackRange          float64 `json:"attackrange"`
		HpRegen              float64 `json:"hpregen"`
		HpRegenPerLevel      float64 `json:"hpregenperlevel"`
		MpRegen              float64 `json:"mpregen"`
		MpRegenPerLevel      float64 `json:"mpregenperlevel"`
		Crit                 float64 `json:"crit"`
		CritPerLevel         float64 `json:"critperlevel"`
		AttackDamage         float64 `json:"attackdamage"`
		AttackDamagePerLevel float64 `json:"attackdamageperlevel"`
		AttackSpeedPerLevel  float64 `json:"attackspeedperlevel"`
		AttackSpeed          float64 `json:"attackspeed"`
	} `json:"stats"`
	Icon   string `json:"icon"`
	Sprite struct {
		URL string  `json:"url"`
		X   float64 `json:"x"`
		Y   float64 `json:"y"`
	} `json:"sprite"`
	Description string `json:"description"`
}

func GetChamps(w http.ResponseWriter, r *http.Request) {
	filePath := `assets\champs.json`
	jsonData, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	w.Write(jsonData)
}

func getChamps() {
	filePath := `assets\champs.json`
	jsonData, err := os.ReadFile(filePath)
	if err != nil {
		log.Println("Error reading JSON file:", err)
		return
	}

	err = json.Unmarshal(jsonData, &champs)
	if err != nil {
		log.Println("Error unmarshaling JSON data:", err)
		return
	}

	log.Println("Champions loaded successfully.")
}

func GetChamp(w http.ResponseWriter, r *http.Request) {
	var champ Champion
	getChamps()
	params := mux.Vars(r)
	name := params["name"]
	champ = *findChampionByID(name)
	if &champ == nil {
		w.Write([]byte("Error"))
		return
	}
	jsonChamp, err := json.Marshal(champ)
	if err != nil {
		w.Write([]byte("No champ"))
		return
	}
	w.Write(jsonChamp)
}

func GetChampParam(w http.ResponseWriter, r *http.Request) {
	var champ Champion
	getChamps()
	params := mux.Vars(r)
	name := params["name"]
	param := params["param"]

	champ = *findChampionByID(name)
	if &champ == nil {
		w.Write([]byte("No such champ in LoL!"))
		return
	}
	statsMap := make(map[string]interface{})
	jsonChamp, err := json.Marshal(champ)
	if err != nil {
		w.Write([]byte("Error encoding JSON"))
		return
	}
	err = json.Unmarshal(jsonChamp, &statsMap)
	if err != nil {
		w.Write([]byte("Error decoding JSON"))
		return
	}
	paramValue, found := statsMap[param]
	if !found {
		w.Write([]byte("Invalid parameter"))
		return
	}

	response := map[string]interface{}{
		param: paramValue,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		w.Write([]byte("Error encoding JSON"))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func findChampionByID(id string) *Champion {
	for _, champ := range champs {
		if champ.ID == id {
			return &champ
		}
	}
	return nil
}
