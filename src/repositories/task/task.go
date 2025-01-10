package task

import (
	taskDTO "github.com/danielsilveiradevbr/api_jira/src/domain/dto/ddsDto"
	"gorm.io/gorm"
)

func SalvaTask(db *gorm.DB, taskDTO *taskDTO.Issues) error {
	// if project, err := project.SalvaProject(db, &taskDTO.Fields.Project); err != nil {
	// 	return err
	// }

	// if sprintModel.Sprint, err = sprint.SalvaSprint(db, taskDTO.Fields.Sprint); err != nil {
	// 	return err
	// }

	// if err = issuetype.SalvaIssueType(db, &taskDTO.Fields.Issuetype); err != nil {
	// 	return err
	// }

	// if err = assignee.SalvaAssigned(db, &taskDTO.Fields.Assinee); err != nil {
	// 	return err
	// }

	// if err = reporter.SalvaReporter(db, &taskDTO.Fields.Reporter); err != nil {
	// 	return err
	// }

	// if err = creator.SalvaCreator(db, &taskDTO.Fields.Creator); err != nil {
	// 	return err
	// }

	// if err = priority.SalvaPriority(db, &taskDTO.Fields.Priority); err != nil {
	// 	return err
	// }

	// if err = status.SalvaStatus(db, &taskDTO.Fields.Status); err != nil {
	// 	return err
	// }

	// if err = resolution.SalvaResolution(db, &taskDTO.Fields.Resolution); err != nil {
	// 	return err
	// }

	// if err = tipoalteracao.SalvatipoAlteracao(db, &taskDTO.Fields.TipoAlteracao); err != nil {
	// 	return err
	// }

	// if err = classificacaoRelevancia.SalvaclassificacaoRelevancia(db, &taskDTO.Fields.ClassificacaoRelevancia); err != nil {
	// 	return err
	// }

	// if err = requerDocumentacao.SalvarequerDocumentacao(db, &taskDTO.Fields.RequerDocumentacao); err != nil {
	// 	return err
	// }

	// if err = requerAnaliseTecnica.SalvaRequerAnaliseTecnica(db, &taskDTO.Fields.RequerAnaliseTecnica); err != nil {
	// 	return err
	// }

	// if err = cliente.SalvaCliente(db, taskDTO.Fields.Cliente); err != nil {
	// 	return err
	// }

	// if err = sku.SalvaSku(db, taskDTO.Fields.SKU); err != nil {
	// 	return err
	// }

	// if err = label.Salvalabel(db, taskDTO.Fields.Labels); err != nil {
	// 	return err
	// }

	// if err = complexidade.SalvaComplexidade(db, &taskDTO.Fields.Complexidade); err != nil {
	// 	return err
	// }
	// task := taskModel.(taskDTO)
	// result := db.First(&task, "codigo = ?", task.Codigo)
	// if result.RowsAffected == 0 {
	// 	res := db.Create(&task)
	// 	if res.Error != nil {
	// 		return res.Error
	// 	}
	// } else {
	// 	db.Save(&task)
	// }
	return nil
}
