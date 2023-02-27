package neoKg

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/neoKg"
	neoKgReq "github.com/flipped-aurora/gin-vue-admin/server/model/neoKg/request"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"log"
	"strconv"
)

type NeoPersonService struct {
}

// CreateBiogMain 创建BiogMain记录
// Author [piexlmax](https://github.com/piexlmax)
func (biogMainService *NeoPersonService) CreateNeoPerson(neoPerson neoKg.NeoPerson) (err error) {
	cypher := "CREATE (p:CPerson {name:$name,personId:$personId, indexYear:$indexYear})"
	var params = map[string]interface{}{
		"personId":  neoPerson.PersonId,
		"name":      neoPerson.Name,
		"indexYear": neoPerson.IndexYear,
	}
	session := global.GVA_NEO4J.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer unsafeClose(session)
	_, err = session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		result, err := tx.Run(cypher, params)
		if err != nil {
			log.Println("write to Neo4j with error:", err)
			return nil, err
		}
		return result.Consume()
	})

	return err
}

// DeleteNeoPerson 删除NeoPerson记录
// Author [piexlmax](https://github.com/piexlmax)
func (neoPersonService *NeoPersonService) DeleteNeoPerson(neoPerson neoKg.NeoPerson) (err error) {
	err = global.MustGetGlobalDBByDBName("known_graph").Unscoped().Delete(&neoPerson).Error
	return err
}

// DeleteNeoPersonByIds 批量删除NeoPerson记录
// Author [piexlmax](https://github.com/piexlmax)
func (neoPersonService *NeoPersonService) DeleteNeoPersonByIds(ids request.IdsReq) (err error) {
	err = global.MustGetGlobalDBByDBName("known_graph").Delete(&[]neoKg.NeoPerson{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateNeoPerson 更新NeoPerson记录
// Author [piexlmax](https://github.com/piexlmax)
func (neoPersonService *NeoPersonService) UpdateNeoPerson(neoPerson neoKg.NeoPerson) (err error) {
	//err = global.GVA_NEO4J.NewSession().Run("known_graph")
	return err
}

// GetNeoPerson 根据id获取NeoPerson记录
// Author [piexlmax](https://github.com/piexlmax)
func (neoPersonService *NeoPersonService) GetNeoPersonById(id int) (neoPerson neoKg.NeoPerson, err error) {
	//query := `MATCH (n:CPerson{personId:$personId})
	//			return n`
	query := `MATCH (n:CPerson) return n`
	var params = map[string]interface{}{"personId": neoPerson.PersonId}
	session := global.GVA_NEO4J.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer unsafeClose(session)
	personResp, err := session.ReadTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		records, err := tx.Run(query, params)
		if err != nil {
			return nil, err
		}
		var result = neoKg.NeoPerson{}
		for records.Next() {
			record := records.Record()
			if node, ok := record.Get("n"); ok {
				person := node.(neo4j.Node)
				personId64, _ := person.Props["personId"]
				personId, _ := strconv.Atoi(strconv.FormatInt(personId64.(int64), 10))
				name, _ := person.Props["name"]
				indexYear64, _ := person.Props["indexYear"]
				indexYear, _ := strconv.Atoi(strconv.FormatInt(indexYear64.(int64), 10))
				result = neoKg.NeoPerson{
					PersonId:  personId,
					Name:      name.(string),
					IndexYear: indexYear,
				}
				break
			}
			//personId, _ := record.Get("personId")
			//name, _ := record.Get("name")
			//indexYear, _ := record.Get("indexYear")
			////a := record.Get("personId").(string)
			//result = neoKg.NeoPerson{
			//	PersonId:  personId.(int),
			//	Name:      name.(string),
			//	IndexYear: indexYear.(int),
			//}
			//break
		}
		return result, nil
	})
	return personResp.(neoKg.NeoPerson), nil
}

// GetNeoPersonInfoList 分页获取NeoPerson记录
// Author [piexlmax](https://github.com/piexlmax)
func (neoPersonService *NeoPersonService) GetNeoPersonInfoList(info neoKgReq.NeoPersonSearch) (list []neoKg.NeoPerson, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.MustGetGlobalDBByDBName("known_graph").Model(&neoKg.NeoPerson{})
	var neoPersons []neoKg.NeoPerson

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&neoPersons).Error
	return neoPersons, total, err
}
