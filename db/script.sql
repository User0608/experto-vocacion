create table usuario(
	username varchar(20) not null,
	password varchar(80) not null,
	estado bool not null default true
);

create table estudiante(
	id serial not null,
	nombre varchar(80) not null,
	apellido_paterno varchar(80) not null,
	apellido_materno varchar(80),
	fecha_nacimiento date not null,
	dni char(8) not null,
	password varchar(80) not null
);

create table casm(
	id serial not null,
	pregunta_a varchar(500),
	pregunta_b varchar(500)
);
create table test(
	id serial not null,
	fecha timestamp default current_timestamp,
	resultado_casm text,
	resultado_berger text,
	resultado_hea text,
	estudiante_id int not null,
	done boolean default false,
	resultado varchar(500)
);
create table test_casm(
	test_id int not null,
	casm_id int not null,
	respuesta_a boolean default false,
	respuesta_b boolean default false	
);

create table berger(
	id serial not null,
	pregunta_a varchar(500) not null,
	pregunta_b varchar(500) not null
);

create table test_berger(
	test_id integer,
	berger_id integer,
	response integer default 0
);
create table hea(
	id serial not null,
	pregunta varchar(500)	
);
create table test_hea(
	test_id integer not null,
	hea_id integer not null,
	respuesta char(1)
);
alter table usuario
	add constraint pk_usuario
		primary key(username);
		
alter table estudiante 
	add constraint fk_estudiante 
		primary key(id);


alter table casm 
	add constraint pk_casm primary key(id);
	
alter table test
	add constraint pk_test primary key(id),
	add constraint fk_test__estudiante
		foreign key(estudiante_id) references estudiante(id);
	
alter table test_casm 
	add constraint pk_test_casm primary key(test_id,casm_id),
	add constraint fk_test_casm__test foreign key(test_id) references test(id),
	add constraint fk_test_casm__casm foreign key(casm_id) references casm(id);
		
alter table berger 
	add constraint pk_berger primary key(id);

alter table test_berger
	add constraint pk_test_berger primary key(test_id,berger_id),
	add constraint fk_test_berger__test foreign key(test_id)
		references test(id),
	add constraint fk_test_berger__berger foreign key(berger_id)
		references berger(id);

alter table hea
	add constraint pk_hea primary key(id);
	
alter table test_hea
	add constraint pk_test_hea primary key(test_id,hea_id),
	add constraint fk_test_hea__test foreign key(test_id) references test(id),
	add constraint fk_test_hea__hea foreign key(hea_id) references hea(id);