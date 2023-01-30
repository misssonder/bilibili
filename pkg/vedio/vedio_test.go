package vedio

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
	bvID, err := ExtractBvID("https://www.bilibili.com/video/BV1kd4y1W7RG/?spm_id_from=333.999.0.0")
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
