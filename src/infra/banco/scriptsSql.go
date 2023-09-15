package banco

func ExexActionsDb() []string {
	var sqls = []string{
		criaTableSprint(),
		criaTableProject(),
		criaTableClientes(),
		criaTableIssueType(),
		criaTablePrioritys(),
		criaTableStatus(),
		craiTableUsers(),
		criaTableTasks(),
		criaTableTasksStatus(),
		criaTableTasksTimers(),
	}
	return sqls
}

func criaTableSprint() string {
	return `
		CREATE TABLE IF NOT EXISTS sprints
		(
			id bigserial NOT NULL,
			id_jira bigint,
			complete_date timestamp,
			start_date timestamp,
			end_date timestamp,
			active_date timestamp,
			data_criacao timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
			data_update timestamp,
			nome character varying(100),
			status character varying(100),
			CONSTRAINT pk_sprints PRIMARY KEY (id)
		);`
}

func criaTableProject() string {
	return `
		CREATE TABLE IF NOT EXISTS projects
		(
			id bigserial NOT NULL,
			id_jira bigint,
			key_jira character varying(200) NOT NULL,
			descricao text,
			data_criacao timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
			data_update timestamp,
			CONSTRAINT pk_projects PRIMARY KEY (id)
		);`
}

func criaTableClientes() string {
	return `
		CREATE TABLE IF NOT EXISTS public.clientes
		(
			id bigserial NOT NULL,
			nome character varying(255) NOT NULL,
			data_update timestamp,
			data_criacao timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
			CONSTRAINT pk_clientes PRIMARY KEY (id)
		);`
}

func criaTableIssueType() string {
	return `
		CREATE TABLE IF NOT EXISTS public.issue_type
		(
			id bigserial NOT NULL,
			nome character varying(200),
			data_criacao timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
			data_update timestamp,
			CONSTRAINT pk_issue_type PRIMARY KEY (id)
		);`
}

func criaTablePrioritys() string {
	return `
	  	CREATE TABLE IF NOT EXISTS public.prioritys
		(
			id bigserial NOT NULL,
			id_jira bigint,
			descricao text,
			data_criacao timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
			data_update timestamp,
			CONSTRAINT pk_prioritys PRIMARY KEY (id)
		);`
}

func criaTableStatus() string {
	return `
		CREATE TABLE IF NOT EXISTS public.status
		(
			id bigserial NOT NULL,
			id_jira bigint,
			descricao text,
			data_criacao timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
			data_update timestamp,
			CONSTRAINT pk_status PRIMARY KEY (id)
		);`
}

func craiTableUsers() string {
	return `
		CREATE TABLE IF NOT EXISTS public.users
		(
			id bigserial NOT NULL,
			nome character varying(200) NOT NULL,
			data_criacao timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
			data_update timestamp,
			key_jira character varying(200),
			email character varying(200),
			avatar text,
			display_name character varying(200),
			CONSTRAINT pk_users PRIMARY KEY (id)
		);`
}

func criaTableTasks() string {
	return `
		CREATE TABLE IF NOT EXISTS public.tasks
		(
			id bigserial NOT NULL,
			id_jira bigint,
			key_jira character varying(200) NOT NULL,
			descricao text,
			data_criacao timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
			data_update timestamp,
			id_user bigint,
			id_report bigint,
			progress_total bigint,
			progress bigint,
			perc_progres numeric(18,4),
			summary text,
			id_project bigint,
			id_priority bigint,
			id_status bigint,
			tipo character varying(60),
			id_issue_type bigint,
			createdat timestamp,
			id_sprint bigint,
			qtde_retrabalho integer,
			id_cliente bigint,
			duplicacao_codigo integer,
			padronizacao_codigo integer,
			documentacao_codigo integer,
			legibilidade_codigo integer,
			simplicidade_codigo integer,
			modularidade_codigo integer,
			target_end timestamp,
			resolution_date timestamp,
			horas_prev numeric,
			complexidade integer,
			CONSTRAINT pk_tasks PRIMARY KEY (id),
			CONSTRAINT fk_priority FOREIGN KEY (id_priority)
				REFERENCES public.prioritys (id),
			CONSTRAINT fk_project FOREIGN KEY (id_project)
				REFERENCES public.projects (id),
			CONSTRAINT fk_report FOREIGN KEY (id_report)
				REFERENCES public.users (id),
			CONSTRAINT fk_status FOREIGN KEY (id_status)
				REFERENCES public.status (id),
			CONSTRAINT fk_task_type FOREIGN KEY (id_issue_type)
				REFERENCES public.issue_type (id),
			CONSTRAINT fk_tasks_cliente FOREIGN KEY (id_cliente)
				REFERENCES public.clientes (id),
			CONSTRAINT fk_tasks_sprints FOREIGN KEY (id_sprint)
				REFERENCES public.sprints (id),
			CONSTRAINT fk_user FOREIGN KEY (id_user)
				REFERENCES public.users (id)
		);`
}

func criaTableTasksStatus() string {
	return `
	CREATE TABLE IF NOT EXISTS public.tasks_status
	(
		id bigserial NOT NULL,
		id_status bigint,
		id_task bigint,
		id_user bigint,
		data_criacao timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
		data_update timestamp,
		CONSTRAINT pk_tasks_status PRIMARY KEY (id),
		CONSTRAINT fk_status_task FOREIGN KEY (id_status)
			REFERENCES public.status (id),
		CONSTRAINT fk_task_status FOREIGN KEY (id_task)
			REFERENCES public.tasks (id),
		CONSTRAINT fk_user_task FOREIGN KEY (id_user)
			REFERENCES public.users (id)
	);`
}

func criaTableTasksTimers() string {
	return `
		CREATE TABLE IF NOT EXISTS public.tasks_timers
		(
			id bigserial NOT NULL,
			id_task bigint,
			id_user bigint,
			tempo integer,
			data_criacao timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
			data_update timestamp,
			id_jira bigint,
			CONSTRAINT pk_tasks_timers PRIMARY KEY (id),
			CONSTRAINT fk_task_timers FOREIGN KEY (id_task)
				REFERENCES public.tasks (id),
			CONSTRAINT fk_user_timers FOREIGN KEY (id_user)
				REFERENCES public.users (id)
		);
	`
}
