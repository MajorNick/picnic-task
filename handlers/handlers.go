package handlers

import (
	//	"fmt"
	"fmt"
	"io"
	"strings"

	"net/http"
	"os"
	"text/template"

	"github.com/MajorNick/picnic-task/parser"

	oset "github.com/MajorNick/orderedSet"
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
)


var tmpl *template.Template

type Data struct {
    Items []parser.Respodent
}
func init(){
	tmpl = template.Must(template.ParseGlob("htmls/*.html"))
}

func Rawdata(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		Sort(w,r)
		return
	}


	
	dataP := parser.GetRawData()
	
	var data Data
	data.Items = *dataP
	tmpl, _ := template.ParseFiles("./htmls/rawdata.html")
	tmpl.Execute(w,data)
	
}

func StartHandler(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w,"index.html",nil)

}
func PieChart(w http.ResponseWriter, r *http.Request){
	
	mp := parser.GetMappedDataOfAnswers()

	items := make([]opts.PieData,0)
	for i,v:=range mp{
		items = append(items, opts.PieData{Name: i,Value: v})
	}
	pie := charts.NewPie()
	pie.SetGlobalOptions(charts.WithTitleOpts(opts.Title{Title:"Most preferable social network"}))
	
	pie.AddSeries("pie", items).
		SetSeriesOptions(
			charts.WithLabelOpts(opts.Label{
				Show:      true,
				Formatter: "{b}: {c}",
			}),
			charts.WithPieChartOpts(opts.PieChart{
				Radius: []string{"40%", "75%"},
			}),
		)
	
	page := components.NewPage()
	page.AddCharts(pie)
	f, err := os.Create("htmls/chart.html")
	if err != nil {
		panic(err)
	}
	page.Render(io.MultiWriter(f))
	tmpl, _ := template.ParseFiles("./htmls/chart.html")
	tmpl.Execute(w,page)
}
func Sort(w http.ResponseWriter, r *http.Request){
	var tmp []parser.Respodent
	var set oset.OrderedSet
	set = parser.GenerateOSet()
	for i:=0;i<set.Size();i++{
		tmp = append(tmp,set.Get(i).(parser.Respodent))
	}
	var data Data
	data.Items = tmp
	tmpl, _ := template.ParseFiles("./htmls/rawdata.html")
	tmpl.Execute(w,data)
}	

func GetValues(w http.ResponseWriter, r *http.Request){
	
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "error in parseForm: %v", err)
		return
	}
	tp := r.FormValue("segment")
	tp = strings.ToLower(tp)
	var tmp []parser.Respodent
	data := parser.GetRawData()
	for _,v := range *data{
		if strings.ToLower(v.SegmentType) == tp{
		tmp = append(tmp,v)
		}
	}
	var dt Data
	dt.Items = tmp
	tmpl, _ := template.ParseFiles("./htmls/filter.html")
	tmpl.Execute(w,dt)

}

