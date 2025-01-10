package fieldsDto

import (
	assineeDto "github.com/danielsilveiradevbr/api_jira/src/domain/dto/ddsDto/assineee"
	classificacaoRelevanciaDto "github.com/danielsilveiradevbr/api_jira/src/domain/dto/ddsDto/classificacaorelevancia"
	complexidadeDto "github.com/danielsilveiradevbr/api_jira/src/domain/dto/ddsDto/complexidade"
	creatorDto "github.com/danielsilveiradevbr/api_jira/src/domain/dto/ddsDto/creator"
	issueTypeDto "github.com/danielsilveiradevbr/api_jira/src/domain/dto/ddsDto/issueType"
	priorityDto "github.com/danielsilveiradevbr/api_jira/src/domain/dto/ddsDto/priority"
	requerAnalisetecnicaDto "github.com/danielsilveiradevbr/api_jira/src/domain/dto/ddsDto/requerAnaliseTecnica"
	requerDocumentacaoDto "github.com/danielsilveiradevbr/api_jira/src/domain/dto/ddsDto/requerDocumentacao"
	resolutionDto "github.com/danielsilveiradevbr/api_jira/src/domain/dto/ddsDto/resolution"
	statusDto "github.com/danielsilveiradevbr/api_jira/src/domain/dto/ddsDto/status"
	tipoAlteracaoDto "github.com/danielsilveiradevbr/api_jira/src/domain/dto/ddsDto/tipoAlteracao"
)

type Fields struct {
	// Customfield13100              any               `json:"customfield_13100"`
	// Customfield13101              any               `json:"customfield_13101"`
	// Customfield14557              any               `json:"customfield_14557"`
	Resolution resolutionDto.Resolution `json:"resolution"`
	// Customfield14308              any               `json:"customfield_14308"`
	// Customfield10622              any               `json:"customfield_10622"`
	// Customfield14309              any               `json:"customfield_14309"`
	// Customfield14306              any               `json:"customfield_14306"`
	// Customfield11718              any               `json:"customfield_11718"`
	// LastViewed                    string            `json:"lastViewed"`
	// Customfield14300              any               `json:"customfield_14300"`
	// Customfield12001              any               `json:"customfield_12001"`
	// Customfield14304              any               `json:"customfield_14304"`
	// Customfield14305              any               `json:"customfield_14305"`
	HorasDev float32  `json:"customfield_14302"`
	Labels   []string `json:"labels"`
	// Customfield10610              any               `json:"customfield_10610"`
	// Customfield10611              any               `json:"customfield_10611"`
	// Customfield12911              any               `json:"customfield_12911"`
	AggregateTimeOriginalEstimate float32 `json:"aggregatetimeoriginalestimate"`
	// Customfield10615              any               `json:"customfield_10615"`
	// Customfield10616              any               `json:"customfield_10616"`
	// Customfield11705              any               `json:"customfield_11705"`
	// Customfield10617              any               `json:"customfield_10617"`
	// Customfield10618              any               `json:"customfield_10618"`
	// Customfield10619              any               `json:"customfield_10619"`
	// Issuelinks                    []any             `json:"issuelinks"`
	Assinee assineeDto.Assignee `json:"assignee"`
	// Components                    []any             `json:"components"`
	// Customfield13200              any               `json:"customfield_13200"`
	// Customfield12901              any               `json:"customfield_12901"`
	// Customfield12900              any               `json:"customfield_12900"`
	// Customfield12902              any               `json:"customfield_12902"`
	// Customfield10604              any               `json:"customfield_10604"`
	// Customfield10605              any               `json:"customfield_10605"`
	// Customfield10606              any               `json:"customfield_10606"`
	// Customfield11818              any               `json:"customfield_11818"`
	// Customfield10607              any               `json:"customfield_10607"`
	// Customfield10608              any               `json:"customfield_10608"`
	// Customfield10609              any               `json:"customfield_10609"`
	// Subtasks                      []any             `json:"subtasks"`
	// Customfield14400              any               `json:"customfield_14400"`
	Reporter resolutionDto.Resolution `json:"reporter"`
	// Customfield13306              any               `json:"customfield_13306"`
	// Customfield13305              any               `json:"customfield_13305"`
	// Customfield10712              any               `json:"customfield_10712"`
	// Progress                      Progress          `json:"progress"`
	// Votes                         Votes             `json:"votes"`
	Issuetype issueTypeDto.Issuetype `json:"issuetype"`
	Project   priorityDto.Priority   `json:"project"`
	// Customfield11001              any               `json:"customfield_11001"`
	// Customfield11002              any               `json:"customfield_11002"`
	// Customfield13302              any               `json:"customfield_13302"`
	// Customfield13301              any               `json:"customfield_13301"`
	// Customfield13303              any               `json:"customfield_13303"`
	// Customfield14504              any               `json:"customfield_14504"`
	// Customfield11910              any               `json:"customfield_11910"`
	ResolutionDate string `json:"resolutiondate"`
	// Watches                       Watches           `json:"watches"`
	// Customfield14500              any               `json:"customfield_14500"`
	// Customfield11909              any               `json:"customfield_11909"`
	Updated string `json:"updated"`
	// Customfield10370              any               `json:"customfield_10370"`
	// Timeoriginalestimate          any               `json:"timeoriginalestimate"`
	Description string `json:"description"`
	// Customfield13400              any               `json:"customfield_13400"`
	// Customfield10800              any               `json:"customfield_10800"`
	Summary string `json:"summary"`
	// Customfield13630              any               `json:"customfield_13630"`
	// Customfield12541              any               `json:"customfield_12541"`
	// Customfield10363              any               `json:"customfield_10363"`
	// Customfield12540              any               `json:"customfield_12540"`
	// Customfield10000              string            `json:"customfield_10000"`
	// Customfield13631              any               `json:"customfield_13631"`
	// Customfield10367              any               `json:"customfield_10367"`
	// Customfield10357              any               `json:"customfield_10357"`
	// Customfield13625              any               `json:"customfield_13625"`
	// Customfield13624              any               `json:"customfield_13624"`
	SKU []string `json:"customfield_13503"`
	// Customfield10359              any               `json:"customfield_10359"`
	// Environment                   any               `json:"environment"`
	// Customfield10910              any               `json:"customfield_10910"`
	// Customfield13626              any               `json:"customfield_13626"`
	// Customfield13629              any               `json:"customfield_13629"`
	// Customfield13628              any               `json:"customfield_13628"`
	// Duedate                       any               `json:"duedate"`
	// FixVersions                   []any             `json:"fixVersions"`
	// Customfield10110              any               `json:"customfield_10110"`
	// Customfield10353              any               `json:"customfield_10353"`
	// Customfield13623              any               `json:"customfield_13623"`
	// Customfield10355              any               `json:"customfield_10355"`
	Cliente []string `json:"customfield_13501"`
	// Customfield12533              any               `json:"customfield_12533"`
	// Customfield10356              any               `json:"customfield_10356"`
	// Customfield13614              any               `json:"customfield_13614"`
	// Customfield10347              any               `json:"customfield_10347"`
	// Customfield13613              any               `json:"customfield_13613"`
	Customfield10105 string `json:"customfield_10105"`
	// Customfield10106              any               `json:"customfield_10106"`
	// Customfield13616              any               `json:"customfield_13616"`
	// Customfield12406              any               `json:"customfield_12406"`
	// Customfield10107              any               `json:"customfield_10107"`
	// Customfield13615              any               `json:"customfield_13615"`
	// Customfield10108              []any             `json:"customfield_10108"`
	// Customfield10901              any               `json:"customfield_10901"`
	// Customfield13617              any               `json:"customfield_13617"`
	// Customfield10109              any               `json:"customfield_10109"`
	// Customfield10902              any               `json:"customfield_10902"`
	// Customfield10903              any               `json:"customfield_10903"`
	// Customfield10904              any               `json:"customfield_10904"`
	// Customfield10906              any               `json:"customfield_10906"`
	// Customfield10907              any               `json:"customfield_10907"`
	// Customfield10908              any               `json:"customfield_10908"`
	// Customfield10909              any               `json:"customfield_10909"`
	// Customfield10340              any               `json:"customfield_10340"`
	// Customfield10341              any               `json:"customfield_10341"`
	// Customfield10342              Customfield10342  `json:"customfield_10342"`
	// Customfield13610              any               `json:"customfield_13610"`
	// Customfield11310              any               `json:"customfield_11310"`
	Priority priorityDto.Priority `json:"priority"`
	Sprint   []string             `json:"customfield_10100"`
	// Customfield12520              any               `json:"customfield_12520"`
	// Customfield10343              any               `json:"customfield_10343"`
	// Customfield10101              string            `json:"customfield_10101"`
	// Customfield13612              any               `json:"customfield_13612"`
	Solucao string `json:"customfield_10344"`
	// Customfield13611              any               `json:"customfield_13611"`
	// Customfield13603              any               `json:"customfield_13603"`
	// Customfield13602              any               `json:"customfield_13602"`
	// Customfield13605              any               `json:"customfield_13605"`
	// Customfield10337              any               `json:"customfield_10337"`
	// Customfield13604              any               `json:"customfield_13604"`
	// Customfield13607              any               `json:"customfield_13607"`
	// Customfield10339              any               `json:"customfield_10339"`
	TimeEstimate float32 `json:"timeestimate"`
	// Customfield13606              any               `json:"customfield_13606"`
	// Customfield11308              any               `json:"customfield_11308"`
	// Versions                      []any             `json:"versions"`
	// Customfield13609              any               `json:"customfield_13609"`
	// Customfield13608              any               `json:"customfield_13608"`
	Status statusDto.Status `json:"status"`
	// Customfield14010              any               `json:"customfield_14010"`
	// Customfield14011              any               `json:"customfield_14011"`
	// Customfield10210              any               `json:"customfield_10210"`
	// Customfield10332              any               `json:"customfield_10332"`
	// Customfield10333              any               `json:"customfield_10333"`
	// Customfield13601              any               `json:"customfield_13601"`
	// Customfield12512              any               `json:"customfield_12512"`
	// Customfield11302              any               `json:"customfield_11302"`
	// Customfield13600              any               `json:"customfield_13600"`
	// Customfield10334              any               `json:"customfield_10334"`
	// Customfield10203              any               `json:"customfield_10203"`
	// Customfield13713              any               `json:"customfield_13713"`
	// Customfield10204              any               `json:"customfield_10204"`
	// Customfield13715              any               `json:"customfield_13715"`
	TargetEnd string `json:"customfield_10205"`
	// Customfield14009              any               `json:"customfield_14009"`
	// Customfield10206              any               `json:"customfield_10206"`
	// Customfield13714              any               `json:"customfield_13714"`
	// Customfield10207              any               `json:"customfield_10207"`
	// Customfield13717              any               `json:"customfield_13717"`
	AggregateTimeEstimate float64 `json:"aggregatetimeestimate"`
	// Customfield10208              any               `json:"customfield_10208"`
	// Customfield13716              any               `json:"customfield_13716"`
	// Customfield10209              any               `json:"customfield_10209"`
	// Customfield13718              any               `json:"customfield_13718"`
	Creator              creatorDto.Creator                           `json:"creator"`
	RequerAnaliseTecnica requerAnalisetecnicaDto.RequerAnaliseTecnica `json:"customfield_14000"`
	// Customfield14003              any               `json:"customfield_14003"`
	// Customfield14004              any               `json:"customfield_14004"`
	// Customfield14001              any               `json:"customfield_14001"`
	// Customfield14002              any               `json:"customfield_14002"`
	// Aggregateprogress             Aggregateprogress `json:"aggregateprogress"`
	// Customfield11410              any               `json:"customfield_11410"`
	// Customfield10321              any               `json:"customfield_10321"`
	// Customfield10200              any               `json:"customfield_10200"`
	// Customfield13711              any               `json:"customfield_13711"`
	// Customfield10322              any               `json:"customfield_10322"`
	// Customfield10201              any               `json:"customfield_10201"`
	// Customfield14005              any               `json:"customfield_14005"`
	// Customfield10323              any               `json:"customfield_10323"`
	// Customfield13710              any               `json:"customfield_13710"`
	// Customfield10202              any               `json:"customfield_10202"`
	// Customfield14006              any               `json:"customfield_14006"`
	// Customfield13702              any               `json:"customfield_13702"`
	Customfield10434 string `json:"customfield_10434"`
	// Customfield11403              any               `json:"customfield_11403"`
	// Customfield13701              any               `json:"customfield_13701"`
	// Customfield10435              any               `json:"customfield_10435"`
	// Customfield13704              any               `json:"customfield_13704"`
	// Customfield10316              any               `json:"customfield_10316"`
	// Customfield13703              any               `json:"customfield_13703"`
	// Customfield13706              any               `json:"customfield_13706"`
	// Customfield13705              any               `json:"customfield_13705"`
	// Customfield13708              any               `json:"customfield_13708"`
	// Customfield13707              any               `json:"customfield_13707"`
	// Customfield11408              any               `json:"customfield_11408"`
	// Customfield13709              any               `json:"customfield_13709"`
	Timespent float64 `json:"timespent"`
	// Customfield10430              any               `json:"customfield_10430"`
	AggregateTimespent float64 `json:"aggregatetimespent"`
	// Customfield10431              any               `json:"customfield_10431"`
	Customfield10432 string `json:"customfield_10432"`
	Customfield10433 string `json:"customfield_10433"`
	// Customfield10423              any               `json:"customfield_10423"`
	// Customfield10424              any               `json:"customfield_10424"`
	// Customfield11636              any               `json:"customfield_11636"`
	// Customfield10426              any               `json:"customfield_10426"`
	// Customfield10427              any               `json:"customfield_10427"`
	// Customfield10428              any               `json:"customfield_10428"`
	// Customfield10429              any               `json:"customfield_10429"`
	// Customfield10308              any               `json:"customfield_10308"`
	Workratio int    `json:"workratio"`
	Created   string `json:"created"`
	// Customfield14100              any               `json:"customfield_14100"`
	// Customfield14101              any               `json:"customfield_14101"`
	// Customfield10420              any               `json:"customfield_10420"`
	// Customfield10421              any               `json:"customfield_10421"`
	ResumoAlteracao string `json:"customfield_13931"`
	// Customfield10301              any               `json:"customfield_10301"`
	RequerDocumentacao requerDocumentacaoDto.RequerDocumentacao `json:"customfield_13930"`
	// Customfield11623              any               `json:"customfield_11623"`
	// Customfield11502              any               `json:"customfield_11502"`
	// Customfield13801              any               `json:"customfield_13801"`
	// Customfield13800              any               `json:"customfield_13800"`
	// Customfield13921              any               `json:"customfield_13921"`
	// Customfield10414              any               `json:"customfield_10414"`
	// Customfield11625              any               `json:"customfield_11625"`
	// Customfield13924              any               `json:"customfield_13924"`
	// Customfield11624              any               `json:"customfield_11624"`
	// Customfield13925              string  `json:"customfield_13925"`
	// Customfield13928              any               `json:"customfield_13928"`
	// Customfield13927              any               `json:"customfield_13927"`
	// Customfield13929              any               `json:"customfield_13929"`
	// Customfield13005              any               `json:"customfield_13005"`
	// Customfield11621              any               `json:"customfield_11621"`
	// Customfield13920              any               `json:"customfield_13920"`
	// Customfield11620              any               `json:"customfield_11620"`
	// Customfield12701              any               `json:"customfield_12701"`
	// Customfield13911              any               `json:"customfield_13911"`
	// Customfield11614              any               `json:"customfield_11614"`
	// Customfield11613              any               `json:"customfield_11613"`
	// Customfield13912              any               `json:"customfield_13912"`
	// Customfield11616              any               `json:"customfield_11616"`
	// Customfield13915              any               `json:"customfield_13915"`
	// Customfield13914              any               `json:"customfield_13914"`
	// Customfield11618              any               `json:"customfield_11618"`
	// Customfield13917              any               `json:"customfield_13917"`
	// Customfield11617              any               `json:"customfield_11617"`
	// Customfield13919              any               `json:"customfield_13919"`
	// Customfield11619              any               `json:"customfield_11619"`
	// Customfield13918              any               `json:"customfield_13918"`
	// Customfield14200              any               `json:"customfield_14200"`
	// Customfield10400              any               `json:"customfield_10400"`
	// Customfield11724              any               `json:"customfield_11724"`
	// Customfield13902              any               `json:"customfield_13902"`
	// Customfield13901              any               `json:"customfield_13901"`
	// Customfield11844              any               `json:"customfield_11844"`
	// Customfield11605              any               `json:"customfield_11605"`
	ClassificacaoRelevancia classificacaoRelevanciaDto.ClassificacaoRelevancia `json:"customfield_13904"`
	// Customfield11604              any               `json:"customfield_11604"`
	// Customfield13903              string            `json:"customfield_13903"`
	TipoAlteracao tipoAlteracaoDto.TipoAlteracao `json:"customfield_13905"`
	// Customfield13908              any               `json:"customfield_13908"`
	// Customfield13907              any               `json:"customfield_13907"`
	HorasTest          float32                      `json:"customfield_14539"`
	PadronizacaoCodigo float32                      `json:"customfield_14549"`
	DuplicacaoCodigo   float32                      `json:"customfield_14548"`
	DocumentacaoCodigo float32                      `json:"customfield_14550"`
	LegibilidadeCodigo float32                      `json:"customfield_14551"`
	SimplicidadeCodigo float32                      `json:"customfield_14552"`
	ModularidadeCodigo float32                      `json:"customfield_14553"`
	QtdeRetrabalho     float32                      `json:"customfield_14544"`
	Complexidade       complexidadeDto.Complexidade `json:"customfield_14559"`
}
