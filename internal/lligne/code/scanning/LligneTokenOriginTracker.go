//
// # Helper facility to track token line and column as a side effect of scanning.
//
// (C) Copyright 2023 Martin E. Nordberg III
// Apache 2.0 License
//

package scanning

//---------------------------------------------------------------------------------------------------------------------

type ILligneTokenOriginTracker interface {
	GetOrigin(sourcePos int) LligneOrigin
}

//---------------------------------------------------------------------------------------------------------------------

type LligneTokenOriginTracker struct {
	fileName         string
	newLinePositions []int
}

//---------------------------------------------------------------------------------------------------------------------

func NewLligneTokenOriginTracker(fileName string) LligneTokenOriginTracker {
	return LligneTokenOriginTracker{
		fileName:         fileName,
		newLinePositions: append(make([]int, 0), -1),
	}
}

//---------------------------------------------------------------------------------------------------------------------

func (t *LligneTokenOriginTracker) AppendNewLinePosition(sourcePos int) {
	t.newLinePositions = append(t.newLinePositions, sourcePos)
}

//---------------------------------------------------------------------------------------------------------------------

func (t *LligneTokenOriginTracker) GetOrigin(sourcePos int) LligneOrigin {

	priorNewLinePos := 0
	if len(t.newLinePositions) > 0 {
		iMin := 0
		iMax := len(t.newLinePositions)
		for iMax-iMin > 1 {
			iMid := (iMin + iMax) / 2
			if sourcePos > t.newLinePositions[iMid] {
				iMin = iMid
			} else {
				iMax = iMid
			}
		}
		priorNewLinePos = iMin
	}

	return LligneOrigin{
		FileName: t.fileName,
		Line:     priorNewLinePos + 1,
		Column:   sourcePos - t.newLinePositions[priorNewLinePos],
	}

}

//---------------------------------------------------------------------------------------------------------------------
