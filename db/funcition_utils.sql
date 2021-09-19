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

-- consulta si un test es eliminable o no..
create or replace function check_test_isnot_used(
	_test_id int
) returns boolean 
as
$$
declare 
		_casm int = 0;
		_berger integer =0;
		_hea int =0;
begin	
	_casm := (select count(tc) from test_casm tc where tc.test_id =_test_id);
	_berger := (select count(tc) from test_casm tc where tc.test_id =_test_id);
	_hea := (select count(th) from test_hea th where th.test_id =_test_id);
	if _casm !=0 or _berger !=0 or _hea  !=0 then 
		return false;
	end if;
	return true;
end
$$
language plpgsql;