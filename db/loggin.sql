create or replace function ps_sign_admin(
	_user_name varchar(60),
	_password varchar(60)
)
returns setof usuario
as
$$
	begin
		return query select *
			from usuario u
			where u.username = _user_name and u.password =_password;
	end;
$$
language plpgsql;



create or replace function ps_sign_in(
	_user_name varchar(60),
	_password varchar(60)
)
returns setof estudiante
as
$$
	begin
		return query select *
			from estudiante e
			where e.dni = _user_name and e.password =_password;
	end;
$$
language plpgsql;
