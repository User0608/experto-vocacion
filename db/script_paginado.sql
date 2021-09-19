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
	state text	
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
		case when tc.test_id isnull then 'none'
		else 'done' end as "state"	
			from casm c		
			left join (select tcc.* from test_casm tcc 
					   where tcc.test_id = _test_id ) tc
			on c.id = tc.casm_id
			order by c.id limit _items offset _items*(_page - 1);
end
$$
language plpgsql;
