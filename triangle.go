/*
read me
This file is for pratcing a custom plugin for telegraf
to help show and create progress


*/



package main


import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"math"
	
	// "github.com/influxdata/telegraf"

	// "github.com/influxdata/telegraf/plugins/inputs"	
)

/*
trig struct
*/
type Trig struct {
x float64
Amplitude float64

}
/*
sample configuration
*/

var TrigConfig = `
## Set the amplitude
amplitude = 10.0
`
/*
sample configuration
*/
func (s *Trig) SampleConfig string {
return TrigConfig
}
/*
sample description to explain the plug in
*/
func (s *Trig) Description string {
	return "returns sine and cosie waves"
	}
/*
gather  configuration function to gather all of the metrics 
*/
func (s *Trig) Gather (acc telegraf.Accumulator) error {
		
	
	sinner := math.Sin((s.x*math.Pi)/5.0) * s.Amplitude
	cosinner := math.Cos((s.x*math.Pi)/5.0) * s.Amplitude

	fields := make(map[string]interface{})
	fields["sine"] = sinner
	fields["cosine"] = cosinner

	tags := make(map[string]string)

	s.x += 1.0
	acc.AddFields("trig", fields, tags)
	
	
	return nil
}	
/*
initial state fuction that intitates the plug in
*/
func init(){
	inputs.Add("trig", func() telegraf.Input { return &Trig{x: 0.0} })
}





