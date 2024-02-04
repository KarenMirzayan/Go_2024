package main

import (
	"net/http"

	"champinfo.com/api/pkg/league"
	"github.com/gorilla/mux"
)

func main() {
	// _GetChamps()
	r := mux.NewRouter()

	r.HandleFunc("/champions", pkg.GetChamps).Methods("GET")
	r.HandleFunc("/champions/{name}", pkg.GetChamp).Methods("GET")
	r.HandleFunc("/champions/{name}/{param}", pkg.GetChampParam).Methods("GET")
	r.HandleFunc("/health-check", pkg.HealthCheck).Methods("GET")
	http.Handle("/", r)

	http.ListenAndServe(":8080", r)
}

// var champs []Champion

// type Champion struct {
// 	ID    string   `json:"id"`
// 	Key   string   `json:"key"`
// 	Name  string   `json:"name"`
// 	Title string   `json:"title"`
// 	Tags  []string `json:"tags"`
// 	Stats struct {
// 		HP                   float64 `json:"hp"`
// 		HpPerLevel           float64 `json:"hpperlevel"`
// 		MP                   float64 `json:"mp"`
// 		MpPerLevel           float64 `json:"mpperlevel"`
// 		MoveSpeed            float64 `json:"movespeed"`
// 		Armor                float64 `json:"armor"`
// 		ArmorPerLevel        float64 `json:"armorperlevel"`
// 		SpellBlock           float64 `json:"spellblock"`
// 		SpellBlockPerLevel   float64 `json:"spellblockperlevel"`
// 		AttackRange          float64 `json:"attackrange"`
// 		HpRegen              float64 `json:"hpregen"`
// 		HpRegenPerLevel      float64 `json:"hpregenperlevel"`
// 		MpRegen              float64 `json:"mpregen"`
// 		MpRegenPerLevel      float64 `json:"mpregenperlevel"`
// 		Crit                 float64 `json:"crit"`
// 		CritPerLevel         float64 `json:"critperlevel"`
// 		AttackDamage         float64 `json:"attackdamage"`
// 		AttackDamagePerLevel float64 `json:"attackdamageperlevel"`
// 		AttackSpeedPerLevel  float64 `json:"attackspeedperlevel"`
// 		AttackSpeed          float64 `json:"attackspeed"`
// 	} `json:"stats"`
// 	Icon   string `json:"icon"`
// 	Sprite struct {
// 		URL string  `json:"url"`
// 		X   float64 `json:"x"`
// 		Y   float64 `json:"y"`
// 	} `json:"sprite"`
// 	Description string `json:"description"`
// }

// func GetChamps(w http.ResponseWriter, r *http.Request) {
// 	filePath := `assets\champs.json`
// 	jsonData, err := os.ReadFile(filePath)
// 	if err != nil {
// 		panic(err)
// 	}
// 	w.Write(jsonData)
// }

// func _GetChamps() {
// 	filePath := `assets\champs.json`
// 	jsonData, err := os.ReadFile(filePath)
// 	if err != nil {
// 		log.Println("Error reading JSON file:", err)
// 		return
// 	}

// 	err = json.Unmarshal(jsonData, &champs)
// 	if err != nil {
// 		log.Println("Error unmarshaling JSON data:", err)
// 		return
// 	}

// 	log.Println("Champions loaded successfully.")
// }

// func GetChamp(w http.ResponseWriter, r *http.Request) {
// 	var champ Champion
// 	params := mux.Vars(r)
// 	name := params["name"]
// 	champ = *findChampionByID(name)
// 	if &champ == nil {
// 		w.Write([]byte("Error"))
// 		return
// 	}
// 	jsonChamp, err := json.Marshal(champ)
// 	if err != nil {
// 		w.Write([]byte("No champ"))
// 		return
// 	}
// 	w.Write(jsonChamp)
// }

// func findChampionByID(id string) *Champion {
// 	for _, champ := range champs {
// 		if champ.ID == id {
// 			return &champ
// 		}
// 	}
// 	return nil
// }
