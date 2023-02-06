package video

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvert(t *testing.T) {
	t.Run("av2bv", func(t *testing.T) {
		bvID := AIDtoBvID(4606803)
		assert.Equal(t, bvID, "BV1gs411B7y4")
	})
	t.Run("bv2av", func(t *testing.T) {
		aid := BvIDToAID("BV1gs411B7y4")
		assert.Equal(t, aid, 4606803)
	})
}

func TestExtractBvID(t *testing.T) {
	bvID, err := ExtractBvID("https://www.bilibili.com/video/BV1sy4y197KP/?spm_id_from=333.337.search-card.all.click&vd_source=76326787bdfce30577382b0e7e18f35c")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(bvID)
	bvID, err = ExtractBvID("BV1kd4y1W7RG")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(bvID)
}

func TestExtractSsID(t *testing.T) {
	t.Log(IsSSID("https://www.bilibili.com/bangumi/play/ss33622?from_spmid=666.24.0.0"))
	bvID, err := ExtractSSID("https://www.bilibili.com/bangumi/play/ss33622?from_spmid=666.24.0.0")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(bvID)
}

func TestExtractSSID(t *testing.T) {
	t.Log(IsSSID("https://www.bilibili.com/bangumi/play/ep729217?from_spmid=666.4.banner.1"))
	id, err := ExtractEpID("https://www.bilibili.com/bangumi/play/ep729217?from_spmid=666.4.banner.1")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(id)
}
