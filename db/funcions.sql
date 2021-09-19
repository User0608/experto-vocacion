create or replace function valid_unique_dni(	
	_dni varchar(60)	
)
returns char(2)
as
$$
	begin	
		if (select count(e.*) from estudiante e where e.dni=_dni) = 0 then
			return  'OK';
		else
			return 'FF';
		end if;
	end;
$$
language plpgsql;

create or replace function try_update_user(	
	_estudiateID int,
	_dni varchar(60)	
)
returns char(2)
as
$$
	begin	
		if (select e.dni from estudiante e where e.id=_estudiateID) = _dni then
			return  'NU';
		else
			return 'UU';
		end if;
	end;
$$
language plpgsql;
