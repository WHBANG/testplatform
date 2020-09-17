package proto

var MeaningConfig *Meaning // 全局配置

type Meaning struct {
	Subjects       Subjects   `json:"subjects"`        // 事件主体的类型
	EventTypes     EventTypes `json:"event_types"`     // 事件的类型
	DiscardReasons Reasons    `json:"discard_reasons"` // 废弃原因
	ProcessReasons Reasons    `json:"process_reasons"` // 处理原因
}

type Subjects map[string]map[string]Subject // [non_motor]["29000"]"行人"
type Subject struct {
	Name string `json:"name"`
}

type EventTypes map[string]map[int]EventType // [non_motor][1]"闯红灯"
type EventType struct {
	Name       string `json:"name"`
	ExportType string `json:"export_type"`
}

type Reasons map[string]map[int]Reason
type Reason struct {
	Name string `json:"name"`
}

func (m *Meaning) GetEventTypeCodes(t string) []int {

	cs := []int{}
	for k := range m.EventTypes[t] {
		cs = append(cs, k)
	}

	return cs
}

func (m *Meaning) IsEventTypeCodesExist(t string, eventType int) bool {

	for k := range m.EventTypes[t] {
		if eventType == k {
			return true
		}
	}

	return false
}

func (m *Meaning) GetSubjectCodes(t string) []string {

	cs := []string{}
	for k := range m.Subjects[t] {
		cs = append(cs, k)
	}

	return cs
}

func (m *Meaning) GetDiscardReasonCodes(t string) []int {

	cs := []int{}
	for k := range m.DiscardReasons[t] {
		cs = append(cs, k)
	}

	return cs
}

func (m *Meaning) GetProcessReasonCodes(t string) []int {

	cs := []int{}
	for k := range m.ProcessReasons[t] {
		cs = append(cs, k)
	}

	return cs
}

func (m *Meaning) MapExportType(t int) string {
	for _, ets := range m.EventTypes {
		et, ok := ets[t]
		if ok {
			return et.ExportType
		}
	}

	return "0"
}

func (m *Meaning) MapEventType(t int) string {
	for _, ets := range m.EventTypes {
		et, ok := ets[t]
		if ok {
			return et.Name
		}
	}

	return "其他"
}

func (m *Meaning) MapSubject(t string) string {
	for _, sjs := range m.Subjects {
		sj, ok := sjs[t]
		if ok {
			return sj.Name
		}
	}

	return "其他"
}

func (m *Meaning) MapDiscardReason(t int) string {
	for _, drs := range m.DiscardReasons {
		dr, ok := drs[t]
		if ok {
			return dr.Name
		}
	}

	return "其他"
}

func (m *Meaning) MapProcessReason(t int) string {
	for _, drs := range m.ProcessReasons {
		dr, ok := drs[t]
		if ok {
			return dr.Name
		}
	}

	return "其他"
}
