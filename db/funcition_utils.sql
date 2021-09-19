-- consulta, si es posible insertar una respues, para una pregunta espesifica

create or replace function fn_consult_insertable_answer(
        _questionTable varchar(30),
        _test_id int,
		_question_id int
) returns bool as
$$
declare
        resultado int;
begin
        execute 'select count(*) from test_' ||
                quote_ident(_questionTable)||
                ' where test_id = '||
                quote_nullable(_test_id)||
				'and '||quote_ident(_questionTable)||'_id = ' ||
				 quote_nullable(_question_id)
				into resultado;
				
		if resultado = 1 then 
			return false;
		else 
			return true;
		end if;
        EXCEPTION
                WHEN  undefined_table then
                        raise exception 'La tabla % no existe!, funcion sql `fn_consult`',tablename;
end;
$$ language plpgsql;