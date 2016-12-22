package handler

import (
	"fmt"
	"vrv/im/service/timestamp"
	"vrv/timestamp/dao"

	_ "github.com/go-sql-driver/mysql"
)

type TimestampImpl struct {
	*VRVHandler
	Single map[string]int64
	List   map[string]int64
	Seq    map[int64]int64
	IncSep int64
}

func NewTimestampImpl() *TimestampImpl {
	return &TimestampImpl{
		Single: make(map[string]int64),
		List:   make(map[string]int64),
		Seq:    make(map[int64]int64),
		IncSep: INC_SEP_A,
	}
}

/**
  根据对象ID，获取对象类型的自增ID
*/
func (t *TimestampImpl) IncSingle(ID int64, type_a1 int32) (int64, error) {
	key := fmt.Sprintf("%d-%d", ID, type_a1)
	value := t.Single[key]
	if value == 0 {
		//到数据库中获取该类型值
		startInc, err := dao.GetITValue(ID, type_a1)
		if err != nil {
			return value, err
		}
		updateInc := startInc + t.IncSep
		if startInc == 0 {
			err = dao.AddITValue(ID, type_a1, updateInc)
		} else {
			err = dao.UpdateITValue(ID, type_a1, updateInc)
		}
		//根据数据库中有没有值，添加或者更新数据库值
		if err != nil {
			return value, err
		}
		value = startInc + 1
	} else {
		if value%t.IncSep == 0 { //超过了增加步长
			updateInc := value + t.IncSep
			//更新数据库值
			err := dao.UpdateITValue(ID, type_a1, updateInc)
			if err != nil {
				return value, err
			}
		}
		value++
	}
	t.Single[key] = value
	return value, nil
}

func (t *TimestampImpl) GetSingle(ID int64, type_a1 int32) (int64, error) {
	key := fmt.Sprintf("%d-%d", ID, type_a1)
	value := t.Single[key]
	if value == 0 {
		//数据库中获取
		dbval, err := dao.GetITValue(ID, type_a1)
		if err != nil {
			return value, err
		}
		if dbval == 0 {
			return value, nil
		}
		//数据库有值返货数据库值
		t.Single[key] = dbval
		//没有值返回默认1
		return dbval, nil
	}
	return value, nil
}

func (t *TimestampImpl) GetBatch(ID int64, type_a1 []int32) (r []int64, err error) {
	return nil, nil
}

func (t *TimestampImpl) IncList(ID int64, type_a1 int32, subID int64) (r int64, err error) {
	return 0, nil
}

func (t *TimestampImpl) IncBatchList(ID int64, type_a1 int32, sublist []int64) (r []int64, err error) {
	return nil, nil
}

func (t *TimestampImpl) DelList(ID int64, type_a1 int32) (err error) {
	return nil
}

func (t *TimestampImpl) GetList(ID int64, type_a1 int32, subID int64) (r int64, err error) {
	return 0, nil
}

func (t *TimestampImpl) GetBatchList(ID int64, type_a1 int32, sublist []int64) (r []int64, err error) {
	return nil, nil
}

func (t *TimestampImpl) GetAllList(ID int64, type_a1 int32) (r map[int64]int64, err error) {
	return nil, nil
}

func (t *TimestampImpl) IncNewID(type_a1 int32) (r int64, err error) {
	return 0, nil
}

func (t *TimestampImpl) GetNewID(type_a1 int32) (r int64, err error) {
	return 0, nil
}

func (t *TimestampImpl) GetAllTimestamp(ID int64, type_a1 []int32, listType []int32) (r *timestamp.TimestampResult_, err error) {
	return nil, nil
}
