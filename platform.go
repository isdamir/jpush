package jpush

type PlatformType string

const (
	IOS      PlatformType = "ios"
	ANDROID  PlatformType = "android"
	WINPHONE PlatformType = "winphone"
)

type Platform struct {
	Object interface{}
}

func AllPlatform() *Platform {
	return &Platform{
		Object: "all",
	}
}

func (this *Platform) Add(os PlatformType) {
	if this.Object == nil {
		s := []string{string(os)}
		this.Object = s
	} else {
		switch this.Object.(type) {
		case []string:
			this.Object = append(this.Object.([]string), string(os))
		default:
		}

	}
}
