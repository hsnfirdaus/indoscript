package penerjemah

import (
	"indoscript/pengurai"
)

func (p *Penerjemah) NodeJika(node *pengurai.NodeJika) (interface{}, *KesalahanPenerjemah) {
	banding, err := p.panggilNodeBoolean(node.Kondisi)
	if err != nil {
		return nil, err
	}

	if banding.Isi {
		hasil, err := p.panggilNode(node.NodeNode)
		if err != nil {
			return nil, err
		}

		return hasil, nil
	}

	for _, kasus := range node.KasusLain {
		banding, err := p.panggilNodeBoolean(kasus.Kondisi)
		if err != nil {
			return nil, err
		}

		if banding.Isi {
			hasil, err := p.panggilNode(kasus.NodeNode)
			if err != nil {
				return nil, err
			}

			return hasil, nil
		}
	}

	if node.LainLain != nil {
		hasil, err := p.panggilNode(node.LainLain)
		if err != nil {
			return nil, err
		}

		return hasil, nil
	}

	return nil, nil
}
