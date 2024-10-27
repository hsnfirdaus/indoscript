package penerjemah

import (
	"indoscript/penerjemah/jenis"
	"indoscript/pengurai"
)

func (p *Penerjemah) nodeAkar(node *pengurai.NodeAkar) (interface{}, *KesalahanPenerjemah) {
	results := make([]interface{}, 0)

	for _, node := range node.NodeNode {
		balikan, apakahBalikan := node.(*pengurai.NodeBalikan)
		if apakahBalikan {
			return p.nodeBalikan(balikan)
		}
		hasil, err := p.panggilNode(node)
		if err != nil {
			return nil, err
		}
		bal, apBal := hasil.(*jenis.Balikan)
		if apBal {
			return bal, nil
		}
		results = append(results, hasil)
	}

	return results, nil
}
