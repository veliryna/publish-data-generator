package lib

import "math/rand"

type BloodPressure struct {
	systolic        float64
	diastolic       float64
	phase           string
	timer           int
	targetSystolic  float64
	targetDiastolic float64
}

func (bp *BloodPressure) UpdateBloodPressure() ([]int32, bool) {
	changed := false
	completed := false

	switch bp.phase {
	case "normal":
		bp.phase = "rising"
		bp.timer = rand.Intn(6) + 10
		bp.targetSystolic = float64(rand.Intn(40) + 130)
		bp.targetDiastolic = float64(rand.Intn(20) + 80)
	case "rising":
		bp.systolic += (bp.targetSystolic - bp.systolic) / float64(bp.timer)
		bp.diastolic += (bp.targetDiastolic - bp.diastolic) / float64(bp.timer)
		bp.timer--
		changed = true
		if bp.timer <= 0 {
			bp.phase = "peak"
			bp.timer = rand.Intn(3) + 10
		}
	case "peak":
		bp.timer--
		changed = true
		if bp.timer <= 0 {
			bp.phase = "falling"
			bp.timer = rand.Intn(6) + 10
		}
	case "falling":
		bp.systolic -= (bp.systolic - 120) / float64(bp.timer)
		bp.diastolic -= (bp.diastolic - 80) / float64(bp.timer)
		bp.timer--
		changed = true
		if bp.timer <= 0 {
			bp.phase = "normal"
			completed = true
		}
	}

	s := int32(bp.systolic + rand.Float64()*2 - 1)
	d := int32(bp.diastolic + rand.Float64()*2 - 1)
	return []int32{s, d}, changed || completed
}

type BloodOxygen struct {
	value  float64
	phase  string
	timer  int
	target float64
}

func (bo *BloodOxygen) UpdateBloodOxygen() ([]int32, bool) {
	changed := false
	completed := false

	switch bo.phase {
	case "normal":
		bo.phase = "falling"
		bo.timer = rand.Intn(3) + 4
		bo.target = float64(rand.Intn(10) + 85)

	case "falling":
		bo.value += (bo.target - bo.value) / float64(bo.timer)
		bo.timer--
		changed = true
		if bo.timer <= 0 {
			bo.phase = "low"
			bo.timer = rand.Intn(3) + 2
		}

	case "low":
		// hold at low level
		bo.timer--
		changed = true
		if bo.timer <= 0 {
			bo.phase = "rising"
			bo.timer = rand.Intn(4) + 10
			bo.target = float64(rand.Intn(4) + 97) // rise to 97–99
		}

	case "rising":
		bo.value += (bo.target - bo.value) / float64(bo.timer)
		bo.timer--
		changed = true
		if bo.timer <= 0 {
			bo.phase = "normal"
			completed = true
		}
	}

	val := int32(bo.value + rand.Float64()*1.5 - 0.75)
	return []int32{val}, changed || completed
}

type Temperature struct {
	value  float64
	phase  string
	timer  int
	target float64
}

func (t *Temperature) UpdateTemperature() ([]int32, bool) {
	changed := false
	completed := false

	switch t.phase {
	case "normal":
		if rand.Float64() < 0.01 {
			t.phase = "rising"
			t.timer = rand.Intn(4) + 10
			t.target = float64(rand.Intn(16) + 385)
		}

	case "rising":
		t.value += (t.target - t.value) / float64(t.timer)
		t.timer--
		changed = true
		if t.timer <= 0 {
			t.phase = "high"
			t.timer = rand.Intn(3) + 10
		}

	case "high":
		t.timer--
		changed = true
		if t.timer <= 0 {
			t.phase = "falling"
			t.timer = rand.Intn(5) + 10
			t.target = float64(rand.Intn(6) + 365)
		}

	case "falling":
		t.value += (t.target - t.value) / float64(t.timer)
		t.timer--
		changed = true
		if t.timer <= 0 {
			t.phase = "normal"
			completed = true
		}
	}

	val := int32(t.value + rand.Float64()*2 - 1)
	return []int32{val}, changed || completed
}
