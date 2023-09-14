package db

func ExexActionsDb() {
	criaTableSprint()
}

func criaTableSprint() {
	tabelaSprint := `
		CREATE TABLE IF NOT EXISTS sprints
		(
			id bigserial NOT NULL,
			id_jira bigint,
			complete_date timestamp without time zone,
			start_date timestamp without time zone,
			end_date timestamp without time zone,
			active_date timestamp without time zone,
			data_criacao timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
			data_update timestamp without time zone,
			nome character varying(100) COLLATE pg_catalog."default",
			status character varying(100) COLLATE pg_catalog."default",
			CONSTRAINT pk_sprints PRIMARY KEY (id)
		)
		
		CREATE INDEX IF NOT EXISTS idx_sprints
			ON sprints (id);`

}
