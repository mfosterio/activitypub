package activitystreams

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

func TestObjectNew(t *testing.T) {
	var testValue = ObjectID("test")
	var testType = ArticleType

	o := ObjectNew(testType)
	o.ID = testValue

	if o.ID != testValue {
		t.Errorf("APObject Id '%v' different than expected '%v'", o.ID, testValue)
	}
	if o.Type != testType {
		t.Errorf("APObject Type '%v' different than expected '%v'", o.Type, testType)
	}

	n := ObjectNew("")
	n.ID = testValue
	if n.ID != testValue {
		t.Errorf("APObject Id '%v' different than expected '%v'", n.ID, testValue)
	}
	if n.Type != ObjectType {
		t.Errorf("APObject Type '%v' different than expected '%v'", n.Type, ObjectType)
	}

}

func TestActivityVocabularyTypes_Contains(t *testing.T) {
	{
		var invalidType ActivityVocabularyType = "RandomType"

		if ActivityTypes.Contains(ActivityType) {
			t.Errorf("Generic Activity Type '%v' should not be valid", ActivityType)
		}
		for _, inValidType := range ObjectTypes {
			if ActivityTypes.Contains(inValidType) {
				t.Errorf("APObject Type '%v' should be invalid", inValidType)
			}
		}
		if ActivityTypes.Contains(invalidType) {
			t.Errorf("Activity Type '%v' should not be valid", invalidType)
		}
		for _, validType := range ActivityTypes {
			if !ActivityTypes.Contains(validType) {
				t.Errorf("Activity Type '%v' should be valid", validType)
			}
		}
	}
	{
		var invalidType ActivityVocabularyType = "RandomType"

		if IntransitiveActivityTypes.Contains(ActivityType) {
			t.Errorf("Generic Activity Type '%v' should not be valid", ActivityType)
		}
		for _, inValidType := range ActivityTypes {
			if IntransitiveActivityTypes.Contains(inValidType) {
				t.Errorf("APObject Type '%v' should be invalid", inValidType)
			}
		}
		if IntransitiveActivityTypes.Contains(invalidType) {
			t.Errorf("Activity Type '%v' should not be valid", invalidType)
		}
		for _, validType := range IntransitiveActivityTypes {
			if !IntransitiveActivityTypes.Contains(validType) {
				t.Errorf("Activity Type '%v' should be valid", validType)
			}
		}
	}
	{
		var invalidType ActivityVocabularyType = "RandomType"

		if ActivityTypes.Contains(ActivityType) {
			t.Errorf("Generic Activity Type '%v' should not be valid", ActivityType)
		}
		for _, inValidType := range CollectionManagementActivityTypes {
			if !CollectionManagementActivityTypes.Contains(inValidType) {
				t.Errorf("APObject Type '%v' should be valid", inValidType)
			}
		}
		if CollectionManagementActivityTypes.Contains(invalidType) {
			t.Errorf("Activity Type '%v' should not be valid", invalidType)
		}
		for _, validType := range ContentManagementActivityTypes {
			if CollectionManagementActivityTypes.Contains(validType) {
				t.Errorf("Activity Type '%v' should not be valid", validType)
			}
		}
		for _, validType := range ReactionsActivityTypes {
			if CollectionManagementActivityTypes.Contains(validType) {
				t.Errorf("Activity Type '%v' should not be valid", validType)
			}
		}
	}

	{
		var invalidType ActivityVocabularyType = "RandomType"

		if ActivityTypes.Contains(ActivityType) {
			t.Errorf("Generic Activity Type '%v' should not be valid", ActivityType)
		}
		for _, inValidType := range ContentManagementActivityTypes {
			if !ContentManagementActivityTypes.Contains(inValidType) {
				t.Errorf("APObject Type '%v' should be valid", inValidType)
			}
		}
		if ContentManagementActivityTypes.Contains(invalidType) {
			t.Errorf("Activity Type '%v' should not be valid", invalidType)
		}
		for _, validType := range CollectionManagementActivityTypes {
			if ContentManagementActivityTypes.Contains(validType) {
				t.Errorf("Activity Type '%v' should not be valid", validType)
			}
		}
		for _, validType := range ReactionsActivityTypes {
			if ContentManagementActivityTypes.Contains(validType) {
				t.Errorf("Activity Type '%v' should not be valid", validType)
			}
		}
	}
	{
		var invalidType ActivityVocabularyType = "RandomType"

		if ReactionsActivityTypes.Contains(ActivityType) {
			t.Errorf("Generic Activity Type '%v' should not be valid", ActivityType)
		}
		for _, inValidType := range ReactionsActivityTypes {
			if !ReactionsActivityTypes.Contains(inValidType) {
				t.Errorf("APObject Type '%v' should be valid", inValidType)
			}
		}
		if ReactionsActivityTypes.Contains(invalidType) {
			t.Errorf("Activity Type '%v' should not be valid", invalidType)
		}
		for _, validType := range CollectionManagementActivityTypes {
			if ReactionsActivityTypes.Contains(validType) {
				t.Errorf("Activity Type '%v' should not be valid", validType)
			}
		}
		for _, validType := range ContentManagementActivityTypes {
			if ReactionsActivityTypes.Contains(validType) {
				t.Errorf("Activity Type '%v' should not be valid", validType)
			}
		}
	}
	{
		for _, validType := range CollectionTypes {
			if !CollectionTypes.Contains(validType) {
				t.Errorf("Generic Type '%#v' should be valid", validType)
			}
		}
	}
	{
		var invalidType ActivityVocabularyType = "RandomType"

		if ActorTypes.Contains(invalidType) {
			t.Errorf("APObject Type '%v' should not be valid", invalidType)
		}
		for _, validType := range ActorTypes {
			if !ActorTypes.Contains(validType) {
				t.Errorf("APObject Type '%v' should be valid", validType)
			}
		}
	}
	{
		for _, validType := range GenericObjectTypes {
			if !GenericObjectTypes.Contains(validType) {
				t.Errorf("Generic Type '%v' should be valid", validType)
			}
		}
	}
	{
		var invalidType ActivityVocabularyType = "RandomType"

		if ObjectTypes.Contains(invalidType) {
			t.Errorf("APObject Type '%v' should not be valid", invalidType)
		}
		for _, validType := range ObjectTypes {
			if !ObjectTypes.Contains(validType) {
				t.Errorf("APObject Type '%v' should be valid", validType)
			}
		}
	}
}

func TestMarshalJSON(t *testing.T) {
	m := NaturalLanguageValues{
		{
			"en", "test",
		},
		{
			"de", "test",
		},
	}
	result, err := m.MarshalJSON()
	if err != nil {
		t.Errorf("Failed marshaling '%v'", err)
	}
	mRes := "{\"de\":\"test\",\"en\":\"test\"}"
	if string(result) != mRes {
		t.Errorf("Different results '%v' vs. '%v'", string(result), mRes)
	}
	//n := NaturalLanguageValuesNew()
	//result, err := n.MarshalJSON()

	s := make(map[LangRef]string)
	s["en"] = "test"
	n1 := NaturalLanguageValues{{
		"en", "test",
	}}
	result1, err1 := n1.MarshalJSON()
	if err1 != nil {
		t.Errorf("Failed marshaling '%v'", err1)
	}
	mRes1 := "\"test\""
	if string(result1) != mRes1 {
		t.Errorf("Different results '%v' vs. '%v'", string(result1), mRes1)
	}
}

func TestNaturalLanguageValue_MarshalJSON(t *testing.T) {
	p := NaturalLanguageValues{
		{
			"en", "the test",
		},
		{
			"fr", "le test",
		},
	}
	js := "{\"en\":\"the test\",\"fr\":\"le test\"}"
	out, err := p.MarshalJSON()

	if err != nil {
		t.Errorf("Error: '%s'", err)
	}
	if js != string(out) {
		t.Errorf("Different marshal result '%s', instead of '%s'", out, js)
	}
	p1 := NaturalLanguageValues{
		{
			"en", "the test",
		},
	}

	out1, err1 := p1.MarshalJSON()

	if err1 != nil {
		t.Errorf("Error: '%s'", err1)
	}
	txt := "\"the test\""
	if txt != string(out1) {
		t.Errorf("Different marshal result '%s', instead of '%s'", out1, txt)
	}
}

func TestObject_IsLink(t *testing.T) {
	o := ObjectNew(ObjectType)
	o.ID = "test"
	if o.IsLink() {
		t.Errorf("%#v should not be a valid link", o.Type)
	}
	m := ObjectNew(AcceptType)
	m.ID = "test"
	if m.IsLink() {
		t.Errorf("%#v should not be a valid link", m.Type)
	}
}

func TestObject_IsObject(t *testing.T) {
	o := ObjectNew(ObjectType)
	o.ID = "test"
	if !o.IsObject() {
		t.Errorf("%#v should be a valid object", o.Type)
	}
	m := ObjectNew(AcceptType)
	m.ID = "test"
	if !m.IsObject() {
		t.Errorf("%#v should be a valid object", m.Type)
	}
}

func TestObjectsArr_Append(t *testing.T) {
	d := make(ItemCollection, 0)

	val := Object{ID: ObjectID("grrr")}

	d.Append(val)

	if len(d) != 1 {
		t.Errorf("Objects array should have exactly an element")
	}
	if !reflect.DeepEqual(d[0], val) {
		t.Errorf("First item in object array does not match %q", val.ID)
	}
}

func TestRecipients(t *testing.T) {
	bob := PersonNew("bob")
	alice := PersonNew("alice")
	foo := OrganizationNew("foo")
	bar := GroupNew("bar")

	first := make(ItemCollection, 0)
	if len(first) != 0 {
		t.Errorf("Objects array should have exactly an element")
	}

	first.Append(bob)
	first.Append(alice)
	first.Append(foo)
	first.Append(bar)
	if len(first) != 4 {
		t.Errorf("Objects array should have exactly 4(four) elements, not %d", len(first))
	}

	first.Append(bar)
	first.Append(alice)
	first.Append(foo)
	first.Append(bob)
	if len(first) != 8 {
		t.Errorf("Objects array should have exactly 8(eight) elements, not %d", len(first))
	}

	ItemCollectionDeduplication(&first)
	if len(first) != 4 {
		t.Errorf("Objects array should have exactly 4(four) elements, not %d", len(first))
	}

	second := make(ItemCollection, 0)
	second.Append(bar)
	second.Append(foo)

	ItemCollectionDeduplication(&first, &second)
	if len(first) != 4 {
		t.Errorf("First Objects array should have exactly 8(eight) elements, not %d", len(first))
	}
	if len(second) != 0 {
		t.Errorf("Second Objects array should have exactly 0(zero) elements, not %d", len(second))
	}

	_, err := ItemCollectionDeduplication(&first, &second, nil)
	if err != nil {
		t.Errorf("Deduplication with empty array failed")
	}
}

func TestNaturalLanguageValue_Get(t *testing.T) {
	testVal := "test"
	a := NaturalLanguageValues{{NilLangRef, testVal}}
	if a.Get(NilLangRef) != testVal {
		t.Errorf("Invalid Get result. Expected %s received %s", testVal, a.Get(NilLangRef))
	}
}

func TestNaturalLanguageValue_Set(t *testing.T) {
	testVal := "test"
	a := NaturalLanguageValues{{NilLangRef, "ana are mere"}}
	err := a.Set(LangRef("en"), testVal)
	if err != nil {
		t.Errorf("Received error when doing Set %s", err)
	}
}

func TestNaturalLanguageValue_Append(t *testing.T) {
	var a NaturalLanguageValues

	if len(a) != 0 {
		t.Errorf("Invalid initialization of %T. Size %d > 0 ", a, len(a))
	}
	langEn := LangRef("en")
	valEn := "random value"

	a.Append(langEn, valEn)
	if len(a) != 1 {
		t.Errorf("Invalid append of one element to %T. Size %d != 1", a, len(a))
	}
	if a.Get(langEn) != valEn {
		t.Errorf("Invalid append of one element to %T. Value of %q not equal to %q, but %q", a, langEn, valEn, a.Get(langEn))
	}
	langDe := LangRef("de")
	valDe := "randomisch"
	a.Append(langDe, valDe)

	if len(a) != 2 {
		t.Errorf("Invalid append of one element to %T. Size %d != 2", a, len(a))
	}
	if a.Get(langEn) != valEn {
		t.Errorf("Invalid append of one element to %T. Value of %q not equal to %q, but %q", a, langEn, valEn, a.Get(langEn))
	}
	if a.Get(langDe) != valDe {
		t.Errorf("Invalid append of one element to %T. Value of %q not equal to %q, but %q", a, langDe, valDe, a.Get(langDe))
	}
}

func TestLangRef_UnmarshalJSON(t *testing.T) {
	lang := "en-US"
	json := `"` + lang + `"`

	var a LangRef
	a.UnmarshalJSON([]byte(json))

	if string(a) != lang {
		t.Errorf("Invalid json unmarshal for %T. Expected %q, found %q", lang, lang, string(a))
	}
}

func TestNaturalLanguageValue_UnmarshalFullObjectJSON(t *testing.T) {
	langEn := "en-US"
	valEn := "random"
	langDe := "de-DE"
	valDe := "zufällig\n"

	//m := make(map[string]string)
	//m[langEn] = valEn
	//m[langDe] = valDe

	json := `{
		"` + langEn + `": "` + valEn + `",
		"` + langDe + `": "` + valDe + `"
	}`

	a := make(NaturalLanguageValues, 0)
	_ = a.Append(LangRef(langEn), valEn)
	_ = a.Append(LangRef(langDe), valDe)
	err := a.UnmarshalJSON([]byte(json))
	if err != nil {
		t.Error(err)
	}
	for lang, val := range a {
		if val.Ref != LangRef(langEn) && val.Ref != LangRef(langDe) {
			t.Errorf("Invalid json unmarshal for %T. Expected lang %q or %q, found %q", a, langEn, langDe, lang)
		}

		if val.Ref == LangRef(langEn) && val.Value != valEn {
			t.Errorf("Invalid json unmarshal for %T. Expected value %q, found %q", a, valEn, val)
		}
		if val.Ref == LangRef(langDe) && val.Value != valDe {
			t.Errorf("Invalid json unmarshal for %T. Expected value %q, found %q", a, valDe, val)
		}
	}
}

func validateEmptyObject(o Object, t *testing.T) {
	if o.ID != "" {
		t.Errorf("Unmarshalled object %T should have empty ID, received %q", o, o.ID)
	}
	if o.Type != "" {
		t.Errorf("Unmarshalled object %T should have empty Type, received %q", o, o.Type)
	}
	if o.AttributedTo != nil {
		t.Errorf("Unmarshalled object %T should have empty AttributedTo, received %q", o, o.AttributedTo)
	}
	if len(o.Name) != 0 {
		t.Errorf("Unmarshalled object %T should have empty Name, received %q", o, o.Name)
	}
	if len(o.Summary) != 0 {
		t.Errorf("Unmarshalled object %T should have empty Summary, received %q", o, o.Summary)
	}
	if len(o.Content) != 0 {
		t.Errorf("Unmarshalled object %T should have empty Content, received %q", o, o.Content)
	}
	if o.URL != nil {
		t.Errorf("Unmarshalled object %T should have empty URL, received %v", o, o.URL)
	}
	if o.Icon != nil {
		t.Errorf("Unmarshalled object %T should have empty Icon, received %v", o, o.Icon)
	}
	if o.Image != nil {
		t.Errorf("Unmarshalled object %T should have empty Image, received %v", o, o.Image)
	}
	if !o.Published.IsZero() {
		t.Errorf("Unmarshalled object %T should have empty Published, received %q", o, o.Published)
	}
	if !o.StartTime.IsZero() {
		t.Errorf("Unmarshalled object %T should have empty StartTime, received %q", o, o.StartTime)
	}
	if !o.Updated.IsZero() {
		t.Errorf("Unmarshalled object %T should have empty Updated, received %q", o, o.Updated)
	}
	if !o.EndTime.IsZero() {
		t.Errorf("Unmarshalled object %T should have empty EndTime, received %q", o, o.EndTime)
	}
	if o.Duration != 0 {
		t.Errorf("Unmarshalled object %T should have empty Duration, received %q", o, o.Duration)
	}
	if len(o.To) > 0 {
		t.Errorf("Unmarshalled object %T should have empty To, received %q", o, o.To)
	}
	if len(o.Bto) > 0 {
		t.Errorf("Unmarshalled object %T should have empty Bto, received %q", o, o.Bto)
	}
	if len(o.CC) > 0 {
		t.Errorf("Unmarshalled object %T should have empty CC, received %q", o, o.CC)
	}
	if len(o.BCC) > 0 {
		t.Errorf("Unmarshalled object %T should have empty BCC, received %q", o, o.BCC)
	}
}

func TestObject_UnmarshalJSON(t *testing.T) {
	o := Object{}

	dataEmpty := []byte("{}")
	o.UnmarshalJSON(dataEmpty)
	validateEmptyObject(o, t)
}

func TestMimeType_UnmarshalJSON(t *testing.T) {
	m := MimeType("")
	dataEmpty := []byte("")

	m.UnmarshalJSON(dataEmpty)
	if m != "" {
		t.Errorf("Unmarshalled object %T should be an empty string, received %q", m, m)
	}
}

func TestLangRefValue_String(t *testing.T) {
	t.Skipf("TODO")
}

func TestLangRefValue_UnmarshalJSON(t *testing.T) {
	t.Skipf("TODO")
}

func TestLangRefValue_UnmarshalText(t *testing.T) {
	t.Skipf("TODO")
}

func TestLangRef_UnmarshalText(t *testing.T) {
	l := LangRef("")
	dataEmpty := []byte("")

	l.UnmarshalText(dataEmpty)
	if l != "" {
		t.Errorf("Unmarshalled object %T should be an empty string, received %q", l, l)
	}
}

func TestObjectID_UnmarshalJSON(t *testing.T) {
	o := ObjectID("")
	dataEmpty := []byte("")

	o.UnmarshalJSON(dataEmpty)
	if o != "" {
		t.Errorf("Unmarshalled object %T should be an empty string, received %q", o, o)
	}
}

func TestNaturalLanguageValue_UnmarshalJSON(t *testing.T) {
	l := LangRef("")
	dataEmpty := []byte("")

	l.UnmarshalJSON(dataEmpty)
	if l != "" {
		t.Errorf("Unmarshalled object %T should be an empty string, received %q", l, l)
	}
}

func TestNaturalLanguageValue_UnmarshalText(t *testing.T) {
	l := LangRef("")
	dataEmpty := []byte("")

	l.UnmarshalText(dataEmpty)
	if l != "" {
		t.Errorf("Unmarshalled object %T should be an empty string, received %q", l, l)
	}
}

func TestObject_GetID(t *testing.T) {
	a := Object{}
	testVal := "crash$"
	a.ID = ObjectID(testVal)
	if string(*a.GetID()) != testVal {
		t.Errorf("%T should return %q, Received %q", a.GetID, testVal, *a.GetID())
	}
}

func TestObject_GetLink(t *testing.T) {
	a := Object{}
	testVal := "crash$"
	a.ID = ObjectID(testVal)
	if string(a.GetLink()) != testVal {
		t.Errorf("%T should return %q, Received %q", a.GetLink, testVal, a.GetLink())
	}
}

func TestObject_GetType(t *testing.T) {
	a := Object{}
	a.Type = ActorType
	if a.GetType() != ActorType {
		t.Errorf("%T should return %q, Received %q", a.GetType, ActorType, a.GetType())
	}
}

func TestNaturalLanguageValue_First(t *testing.T) {
	t.Skipf("TODO")
}

func TestNaturalLanguageValueNew(t *testing.T) {
	n := NaturalLanguageValuesNew()

	if len(n) != 0 {
		t.Errorf("Initial %T should have length 0, received %d", n, len(n))
	}
}

func TestNaturalLanguageValue_MarshalText(t *testing.T) {
	t.Skipf("TODO")
}

func TestNaturalLanguageValues_Append(t *testing.T) {
	t.Skipf("TODO")
}

func TestNaturalLanguageValues_First(t *testing.T) {
	t.Skipf("TODO")
}

func TestNaturalLanguageValues_Get(t *testing.T) {
	t.Skipf("TODO")
}

func TestNaturalLanguageValues_MarshalJSON(t *testing.T) {
	t.Skipf("TODO")
}

func TestNaturalLanguageValues_MarshalText(t *testing.T) {
	t.Skipf("TODO")
}

func TestNaturalLanguageValues_Set(t *testing.T) {
	t.Skipf("TODO")
}

func TestNaturalLanguageValues_UnmarshalJSON(t *testing.T) {
	{
		lang := []byte{'e', 'n'}
		val := []byte{'a', 'n', 'a', ' ', 'a', 'r', 'e', ' ', 'm', 'e', 'r', 'e', '\n'}
		js := fmt.Sprintf(`[{"%s": "%s"}]`, lang, val)
		n := NaturalLanguageValues{}
		err := n.UnmarshalJSON([]byte(js))
		if err != nil {
			t.Errorf("Unexpected error when unmarshalling %T: %s", n, err)
		}

		if n.Count() != 1 {
			t.Errorf("Invalid number of elements %d, expected %d", n.Count(), 1)
		}
		l := n.First()
		if l.Value != "ana are mere\n" {
			t.Errorf("Invalid %T value %q, expected %q", l, l.Value, "ana are mere\n")
		}
		if l.Ref != "en" {
			t.Errorf("Invalid %T ref %q, expected %q", l, l.Ref, "en")
		}
	}
	{
		ob := make(map[string]string)
		ob["en"] = "ana are mere\n"
		js, err := json.Marshal(ob)
		if err != nil {
			t.Errorf("Unexpected error when marshalling %T: %s", ob, err)
		}
		n := NaturalLanguageValues{}
		err = n.UnmarshalJSON(js)
		if err != nil {
			t.Errorf("Unexpected error when unmarshalling %T: %s", n, err)
		}

		if n.Count() != 1 {
			t.Errorf("Invalid number of elements %d, expected %d", n.Count(), 1)
		}
		l := n.First()
		if l.Value != "ana are mere\n" {
			t.Errorf("Invalid %T value %q, expected %q", l, l.Value, "ana are mere\n")
		}
		if l.Ref != "en" {
			t.Errorf("Invalid %T ref %q, expected %q", l, l.Ref, "en")
		}
	}
}

func TestNaturalLanguageValues_UnmarshalText(t *testing.T) {
	t.Skipf("TODO")
}

func TestNaturalLanguageValuesNew(t *testing.T) {
	t.Skipf("TODO")
}

func TestToObject(t *testing.T) {
	var it Item
	ob := ObjectNew(ArticleType)
	it = ob

	o, err := ToObject(it)
	if err != nil {
		t.Error(err)
	}
	if o != ob {
		t.Errorf("Invalid activity returned by ToObject #%v", ob)
	}

	act := ActivityNew(ObjectID("test"), CreateType, nil)
	it = act

	a, err := ToObject(it)
	if err == nil {
		t.Errorf("Error returned when calling ToObject with activity should not be nil")
	}
	if a != nil {
		t.Errorf("Invalid return by ToObject #%v, should have been nil", a)
	}
}

func TestFlattenObjectProperties(t *testing.T) {
	t.Skipf("TODO")
}

func TestFlattenProperties(t *testing.T) {
	t.Skipf("TODO")
}

func TestTombstone_UnmarshalJSON(t *testing.T) {
	t.Skipf("TODO")
}

func TestToTombstone(t *testing.T) {
	t.Skipf("TODO")
}

func TestToPlace(t *testing.T) {
	t.Skipf("TODO")
}

func TestToProfile(t *testing.T) {
	t.Skipf("TODO")
}

func TestToRelationship(t *testing.T) {
	t.Skipf("TODO")
}

func TestObject_Recipients(t *testing.T) {
	t.Skipf("TODO")
}

func TestPlace_Recipients(t *testing.T) {
	t.Skipf("TODO")
}

func TestProfile_Recipients(t *testing.T) {
	t.Skipf("TODO")
}

func TestRelationship_Recipients(t *testing.T) {
	t.Skipf("TODO")
}

func TestTombstone_Recipients(t *testing.T) {
	t.Skipf("TODO")
}

func TestNaturalLanguageValues_String(t *testing.T) {
	t.Skipf("TODO")
}

func TestNaturalLanguageValues_Count(t *testing.T) {
	t.Skipf("TODO")
}

func TestItemCollectionDeduplication(t *testing.T) {
	t.Skipf("TODO")
}
