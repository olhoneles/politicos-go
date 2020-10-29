// Copyright (c) 2020, Marcelo Jorge Vieira
// Licensed under the AGPL-3.0+ License

package politicos

type Cand2016 struct {
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

func (c Cand2016) GetYear() int {
	return 2016
}

func (c *Cand2016) New(record []string) {
	c.DT_GERACAO = record[0]
	c.HH_GERACAO = record[1]
	c.ANO_ELEICAO = record[2]
	c.CD_TIPO_ELEICAO = record[3]
	c.NM_TIPO_ELEICAO = record[4]
	c.NR_TURNO = record[5]
	c.CD_ELEICAO = record[6]
	c.DS_ELEICAO = record[7]
	c.DT_ELEICAO = record[8]
	c.TP_ABRANGENCIA = record[9]
	c.SG_UF = record[10]
	c.SG_UE = record[11]
	c.NM_UE = record[12]
	c.CD_CARGO = record[13]
	c.DS_CARGO = record[14]
	c.SQ_CANDIDATO = record[15]
	c.NR_CANDIDATO = record[16]
	c.NM_CANDIDATO = record[17]
	c.NM_URNA_CANDIDATO = record[18]
	c.NM_SOCIAL_CANDIDATO = record[19]
	c.NR_CPF_CANDIDATO = record[20]
	c.NM_EMAIL = record[21]
	c.CD_SITUACAO_CANDIDATURA = record[22]
	c.DS_SITUACAO_CANDIDATURA = record[23]
	c.CD_DETALHE_SITUACAO_CAND = record[24]
	c.DS_DETALHE_SITUACAO_CAND = record[25]
	c.TP_AGREMIACAO = record[26]
	c.NR_PARTIDO = record[27]
	c.SG_PARTIDO = record[28]
	c.NM_PARTIDO = record[29]
	c.SQ_COLIGACAO = record[30]
	c.NM_COLIGACAO = record[31]
	c.DS_COMPOSICAO_COLIGACAO = record[32]
	c.CD_NACIONALIDADE = record[33]
	c.DS_NACIONALIDADE = record[34]
	c.SG_UF_NASCIMENTO = record[35]
	c.CD_MUNICIPIO_NASCIMENTO = record[36]
	c.NM_MUNICIPIO_NASCIMENTO = record[37]
	c.DT_NASCIMENTO = record[38]
	c.NR_IDADE_DATA_POSSE = record[39]
	c.NR_TITULO_ELEITORAL_CANDIDATO = record[40]
	c.CD_GENERO = record[41]
	c.DS_GENERO = record[42]
	c.CD_GRAU_INSTRUCAO = record[43]
	c.DS_GRAU_INSTRUCAO = record[44]
	c.CD_ESTADO_CIVIL = record[45]
	c.DS_ESTADO_CIVIL = record[46]
	c.CD_COR_RACA = record[47]
	c.DS_COR_RACA = record[48]
	c.CD_OCUPACAO = record[49]
	c.DS_OCUPACAO = record[50]
	c.VR_DESPESA_MAX_CAMPANHA = record[51]
	c.CD_SIT_TOT_TURNO = record[52]
	c.DS_SIT_TOT_TURNO = record[53]
	c.ST_REELEICAO = record[54]
	c.ST_DECLARAR_BENS = record[55]
	c.NR_PROTOCOLO_CANDIDATURA = record[56]
	c.NR_PROCESSO = record[57]
	c.CD_SITUACAO_CANDIDATO_PLEITO = record[58]
	c.DS_SITUACAO_CANDIDATO_PLEITO = record[59]
	c.CD_SITUACAO_CANDIDATO_URNA = record[60]
	c.DS_SITUACAO_CANDIDATO_URNA = record[61]
	c.ST_CANDIDATO_INSERIDO_URNA = record[62]
}

func (c Cand2016) GetCollectionName() string {
	return "candidatures"
}
