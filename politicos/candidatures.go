// Copyright (c) 2020, Marcelo Jorge Vieira
// Licensed under the AGPL-3.0+ License

package politicos

// FIXME
type Candidatures struct {
	DT_GERACAO                    string
	HH_GERACAO                    string
	ANO_ELEICAO                   string
	CD_TIPO_ELEICAO               string
	NM_TIPO_ELEICAO               string
	NR_TURNO                      string
	CD_ELEICAO                    string
	DS_ELEICAO                    string
	DT_ELEICAO                    string
	TP_ABRANGENCIA                string
	SG_UF                         string
	SG_UE                         string
	NM_UE                         string
	CD_CARGO                      string
	DS_CARGO                      string
	SQ_CANDIDATO                  string
	NR_CANDIDATO                  string
	NM_CANDIDATO                  string
	NM_URNA_CANDIDATO             string
	NM_SOCIAL_CANDIDATO           string
	NR_CPF_CANDIDATO              string
	NM_EMAIL                      string
	CD_SITUACAO_CANDIDATURA       string
	DS_SITUACAO_CANDIDATURA       string
	CD_DETALHE_SITUACAO_CAND      string
	DS_DETALHE_SITUACAO_CAND      string
	TP_AGREMIACAO                 string
	NR_PARTIDO                    string
	SG_PARTIDO                    string
	NM_PARTIDO                    string
	SQ_COLIGACAO                  string
	NM_COLIGACAO                  string
	DS_COMPOSICAO_COLIGACAO       string
	CD_NACIONALIDADE              string
	DS_NACIONALIDADE              string
	SG_UF_NASCIMENTO              string
	CD_MUNICIPIO_NASCIMENTO       string
	NM_MUNICIPIO_NASCIMENTO       string
	DT_NASCIMENTO                 string
	NR_IDADE_DATA_POSSE           string
	NR_TITULO_ELEITORAL_CANDIDATO string
	CD_GENERO                     string
	DS_GENERO                     string
	CD_GRAU_INSTRUCAO             string
	DS_GRAU_INSTRUCAO             string
	CD_ESTADO_CIVIL               string
	DS_ESTADO_CIVIL               string
	CD_COR_RACA                   string
	DS_COR_RACA                   string
	CD_OCUPACAO                   string
	DS_OCUPACAO                   string
	VR_DESPESA_MAX_CAMPANHA       string
	CD_SIT_TOT_TURNO              string
	DS_SIT_TOT_TURNO              string
	ST_REELEICAO                  string
	ST_DECLARAR_BENS              string
	NR_PROTOCOLO_CANDIDATURA      string
	NR_PROCESSO                   string
	CD_SITUACAO_CANDIDATO_PLEITO  string
	DS_SITUACAO_CANDIDATO_PLEITO  string
	CD_SITUACAO_CANDIDATO_URNA    string
	DS_SITUACAO_CANDIDATO_URNA    string
	ST_CANDIDATO_INSERIDO_URNA    string
}

func (c Candidatures) GetCollectionName() string {
	return "candidatures"
}

func (c Candidatures) Cast() Queryable {
	return &c
}
