package domain

type User struct {
	Id       int
	Name     string
	Surname  string
	Email    string
	Password string
}

type Users []User

type Patient struct {
	Id               int
	PatientId        string
	PatientName      string
	PatientSex       string
	PatientAge       string
	PatientBirthdate string
}

type Patients []Patient

type Study struct {
	Id                     int
	StudyId                string
	StudyInstanceUid       string
	StudyDate              string
	StudyTime              string
	StudyDescription       string
	AccessionNumber        string
	ReferringPhysicianName string
}

type Studies []Study

type Series struct {
	Id                int
	SeriesInstanceUID string
	SeriesDate        string
	SeriesTime        string
	SeriesDescription string
	Modality          string
	SeriesNumber      string
	PatientPosition   string
	BodyPartExamined  string
}
type SeriesList []Series

type Instance struct {
	Id                        int
	SOPInstanceUID            string
	InstanceNumber            string
	Rows                      int
	Columns                   int
	ColorSpace                string
	PhotometricInterpretation string
	BitsAllocated             byte
	BitsStored                byte
	SmallestImagePixelValue   int
	LargestImagePixelValue    int
	PixelAspectRation         string
	PixelSpacing              string
}

type Instances []Instance
