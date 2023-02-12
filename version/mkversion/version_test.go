package version

import (
	"strings"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

func mkInfo(gitCommit, otherCommit string, timestamp time.Time, major, minor, patch, changeCount int) verInfo {
	return verInfo{
		major:       major,
		minor:       minor,
		patch:       patch,
		changeCount: changeCount,
		commit:      gitCommit,
		otherCommit: otherCommit,
		timestamp:   timestamp,
	}
}

func TestMkversion(t *testing.T) {
	corpDate := time.Date(2023, time.January, 27, 1, 2, 3, 4, time.UTC)

	tests := []struct {
		in   verInfo
		want string
	}{
		{mkInfo("abcdef", "", corpDate, 0, 98, 0, 0), `
           VERSION_SHORT="0.98.0"
           VERSION_LONG="0.98.0-tabcdef"
           VERSION_GIT_HASH="abcdef"
           VERSION_EXTRA_HASH=""
           VERSION_XCODE="100.98.0"
           VERSION_XCODE_MACOS="273.27.3723"
           VERSION_WINRES="0,98,0,0"
           VERSION_TRACK="stable"
           VERSION_MSIPRODUCT_AMD64="C653B075-AD91-5265-9DF8-0087D35D148D"
           VERSION_MSIPRODUCT_ARM64="1C41380B-A742-5A3C-AF5D-DF7894DD0FB8"
           VERSION_MSIPRODUCT_X86="4ABDDA14-7499-5C2E-A62A-DD435C50C4CB"
           VERSION_COPYRIGHT="Placeholder"
           VERSION_CAPABILITY=42`},
		{mkInfo("abcdef", "", corpDate, 0, 98, 1, 0), `
           VERSION_SHORT="0.98.1"
           VERSION_LONG="0.98.1-tabcdef"
           VERSION_GIT_HASH="abcdef"
           VERSION_EXTRA_HASH=""
           VERSION_XCODE="100.98.1"
           VERSION_XCODE_MACOS="273.27.3723"
           VERSION_WINRES="0,98,1,0"
           VERSION_TRACK="stable"
           VERSION_MSIPRODUCT_AMD64="DFD6DCF2-06D8-5D19-BDA0-FAF31E44EC23"
           VERSION_MSIPRODUCT_ARM64="A4CCF19C-372B-5007-AFD8-1AF661DFF670"
           VERSION_MSIPRODUCT_X86="FF12E937-DDC4-5868-9B63-D35B2050D4EA"
           VERSION_COPYRIGHT="Placeholder"
           VERSION_CAPABILITY=42`},
		{mkInfo("abcdef", "", corpDate, 1, 2, 9, 0), `
           VERSION_SHORT="1.2.9"
           VERSION_LONG="1.2.9-tabcdef"
           VERSION_GIT_HASH="abcdef"
           VERSION_EXTRA_HASH=""
           VERSION_XCODE="101.2.9"
           VERSION_XCODE_MACOS="273.27.3723"
           VERSION_WINRES="1,2,9,0"
           VERSION_TRACK="stable"
           VERSION_MSIPRODUCT_AMD64="D47B5157-FF26-5A10-A94E-50E4529303A9"
           VERSION_MSIPRODUCT_ARM64="91D16F75-2A12-5E12-820A-67B89BF858E7"
           VERSION_MSIPRODUCT_X86="8F1AC1C6-B93B-5C70-802E-6AE9591FA0D6"
           VERSION_COPYRIGHT="Placeholder"
           VERSION_CAPABILITY=42`},
		{mkInfo("abcdef", "", corpDate, 1, 15, 0, 129), `
           VERSION_SHORT="1.15.129"
           VERSION_LONG="1.15.129-tabcdef"
           VERSION_GIT_HASH="abcdef"
           VERSION_EXTRA_HASH=""
           VERSION_XCODE="101.15.129"
           VERSION_XCODE_MACOS="273.27.3723"
           VERSION_WINRES="1,15,129,0"
           VERSION_TRACK="unstable"
           VERSION_MSIPRODUCT_AMD64="89C96952-1FB8-5A4D-B02E-16A8060C56AA"
           VERSION_MSIPRODUCT_ARM64="DB1A2E86-66C4-5CEC-8F4C-7DB805370F3A"
           VERSION_MSIPRODUCT_X86="DC57C0C3-5164-5C92-86B3-2800CEFF0540"
           VERSION_COPYRIGHT="Placeholder"
           VERSION_CAPABILITY=42`},
		{mkInfo("abcdef", "", corpDate, 1, 2, 0, 17), `
           VERSION_SHORT="1.2.0"
           VERSION_LONG="1.2.0-17-tabcdef"
           VERSION_GIT_HASH="abcdef"
           VERSION_EXTRA_HASH=""
           VERSION_XCODE="101.2.0"
           VERSION_XCODE_MACOS="273.27.3723"
           VERSION_WINRES="1,2,0,0"
           VERSION_TRACK="stable"
           VERSION_MSIPRODUCT_AMD64="0F9709AE-0E5E-51AF-BCCD-A25314B4CE8B"
           VERSION_MSIPRODUCT_ARM64="39D5D46E-E644-5C80-9EF8-224AC1AD5969"
           VERSION_MSIPRODUCT_X86="4487136B-2D11-5E42-BD80-B8529F3326F4"
           VERSION_COPYRIGHT="Placeholder"
           VERSION_CAPABILITY=42`},
		{mkInfo("abcdef", "defghi", corpDate, 1, 15, 0, 129), `
           VERSION_SHORT="1.15.129"
           VERSION_LONG="1.15.129-tabcdef-gdefghi"
           VERSION_GIT_HASH="abcdef"
           VERSION_EXTRA_HASH="defghi"
           VERSION_XCODE="101.15.129"
           VERSION_XCODE_MACOS="273.27.3723"
           VERSION_WINRES="1,15,129,0"
           VERSION_TRACK="unstable"
           VERSION_MSIPRODUCT_AMD64="89C96952-1FB8-5A4D-B02E-16A8060C56AA"
           VERSION_MSIPRODUCT_ARM64="DB1A2E86-66C4-5CEC-8F4C-7DB805370F3A"
           VERSION_MSIPRODUCT_X86="DC57C0C3-5164-5C92-86B3-2800CEFF0540"
           VERSION_COPYRIGHT="Placeholder"
           VERSION_CAPABILITY=42`},
		{mkInfo("abcdef", "", corpDate, 1, 2, 0, 17), `
           VERSION_SHORT="1.2.0"
           VERSION_LONG="1.2.0-17-tabcdef"
           VERSION_GIT_HASH="abcdef"
           VERSION_EXTRA_HASH=""
           VERSION_XCODE="101.2.0"
           VERSION_XCODE_MACOS="273.27.3723"
           VERSION_WINRES="1,2,0,0"
           VERSION_TRACK="stable"
           VERSION_MSIPRODUCT_AMD64="0F9709AE-0E5E-51AF-BCCD-A25314B4CE8B"
           VERSION_MSIPRODUCT_ARM64="39D5D46E-E644-5C80-9EF8-224AC1AD5969"
           VERSION_MSIPRODUCT_X86="4487136B-2D11-5E42-BD80-B8529F3326F4"
           VERSION_COPYRIGHT="Placeholder"
           VERSION_CAPABILITY=42`},
		{mkInfo("abcdef", "defghi", corpDate, 1, 15, 0, 129), `
           VERSION_SHORT="1.15.129"
           VERSION_LONG="1.15.129-tabcdef-gdefghi"
           VERSION_GIT_HASH="abcdef"
           VERSION_EXTRA_HASH="defghi"
           VERSION_XCODE="101.15.129"
           VERSION_XCODE_MACOS="273.27.3723"
           VERSION_WINRES="1,15,129,0"
           VERSION_TRACK="unstable"
           VERSION_MSIPRODUCT_AMD64="89C96952-1FB8-5A4D-B02E-16A8060C56AA"
           VERSION_MSIPRODUCT_ARM64="DB1A2E86-66C4-5CEC-8F4C-7DB805370F3A"
           VERSION_MSIPRODUCT_X86="DC57C0C3-5164-5C92-86B3-2800CEFF0540"
           VERSION_COPYRIGHT="Placeholder"
           VERSION_CAPABILITY=42`},
		{mkInfo("abcdef", "", corpDate, 0, 99, 5, 0), ""},            // unstable, patch number not allowed
		{mkInfo("abcdef", "", corpDate, 0, 99, 5, 123), ""},          // unstable, patch number not allowed
		{mkInfo("abcdef", "defghi", time.Time{}, 1, 15, 0, 129), ""}, // missing corpDate
	}

	for _, test := range tests {
		want := strings.ReplaceAll(strings.TrimSpace(test.want), " ", "")
		info, err := mkOutput(test.in)
		if err != nil {
			if test.want != "" {
				t.Errorf("%#v got unexpected error %v", test.in, err)
			}
			continue
		}

		// Normalize some parts so the test outputs above aren't brittle.
		if info.CapabilityVersion == 0 {
			t.Errorf("info.CapabilityVersion not set")
		}
		info.CapabilityVersion = 42

		if !strings.Contains(info.Copyright, "Copyright") || !strings.Contains(info.Copyright, "Tailscale") {
			t.Errorf("info.Copyright not correct, got %q", info.Copyright)
		}
		info.Copyright = "Placeholder"

		got := strings.TrimSpace(info.String())
		if diff := cmp.Diff(got, want); want != "" && diff != "" {
			t.Errorf("%#v wrong output (-got+want):\n%s", test.in, diff)
		}
	}
}
