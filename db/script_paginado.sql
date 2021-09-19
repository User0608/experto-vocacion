-- paginado de preguntas casm

create or replace function preguntas_casm_page(
	_test_id integer,
	_items integer,
	_page  integer	
) returns table(
	id integer,
	question_a varchar(500),
	question_b varchar(500),
	answer_a boolean,
	answer_b boolean,
	done bool	
) as
$$
begin
	if (select count(t) from test t where t.id = _test_id) = 0 then
		return;
	end if;
	return query select
		c.id as id,
		c.pregunta_a as question_a,
		c.pregunta_b as question_b,
		coalesce(tc.respuesta_a,false) as answer_a,
		coalesce(tc.respuesta_b,false) as answer_b,
		case when tc.test_id isnull then false
		else true end as done
			from casm c		
			left join (select tcc.* from test_casm tcc 
					   where tcc.test_id = _test_id ) tc
			on c.id = tc.casm_id
			order by c.id limit _items offset _items*(_page - 1);
end
$$
language plpgsql;

-- pagina para berger
create or replace function preguntas_berger_page(
	_test_id integer,_items integer,_page  integer	
) returns table(
	id integer,
	question_a varchar(500),
	question_b varchar(500),
	answer integer,	
	done bool	
) as
$$
begin
	if (select count(t) from test t where t.id = _test_id) = 0 then
		return;
	end if;
		return query select
		b.id as id,
		b.pregunta_a as question_a,
		b.pregunta_b as question_b,		
		case when tb.response isnull then 0 else tb.response end as answer,
		case when tb.test_id isnull then false else true end as done
			from berger b		
			left join (select tbb.* from test_berger tbb 
					   where tbb.test_id = _test_id ) tb
			on b.id = tb.berger_id
			order by b.id limit _items offset _items*(_page - 1);
end
$$
language plpgsql;

-- paginado para HEA
create or replace function preguntas_hea_page(
	_test_id integer,_items integer,_page  integer	
) returns table(
	id integer,
	question varchar(500),
	answer char(1),	
	done bool	
) as
$$
begin
	if (select count(t) from test t where t.id = _test_id) = 0 then
		return;
	end if;
		return query select
		h.id as id,
		h.pregunta as question,
		case when th.respuesta isnull then '-' else th.respuesta end as answer,
		case when th.test_id isnull then false else true end as done
			from hea h		
			left join (select thh.* from test_hea thh 
					   where thh.test_id = _test_id ) th
			on h.id = th.hea_id
			order by h.id limit _items offset _items*(_page - 1);
end
$$
language plpgsql;