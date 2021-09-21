personalidad(X) :- 
    emotivo('EMOTIVO'),
	activo('ACTIVO'),
	orden('PRIMARIO'),
	X = 'COLERICO'.

personalidad(X) :- 
    emotivo('EMOTIVO'),
	activo('ACTIVO'),
	orden('SECUNDARIO'),
	X = 'PASIONAL'.

personalidad(X) :- 
    emotivo('EMOTIVO'),
	activo('NO ACTIVO'),
	orden('PRIMARIO'),
	X = 'NERVIOSO'.

personalidad(X) :- 
    emotivo('EMOTIVO'),
	activo('NO ACTIVO'),
	orden('SECUNDARIO'),
	X = 'SENTIMENTAL'.

personalidad(X) :- 
    emotivo('NO EMOTIVO'),
	activo('ACTIVO'),
	orden('PRIMARIO'),
	X = 'SANGUINEO'.

personalidad(X) :- 
    emotivo('NO EMOTIVO'),
	activo('ACTIVO'),
	orden('SECUNDARIO'),
	X = 'FLEMATICO'.

personalidad(X) :- 
    emotivo('NO EMOTIVO'),
	activo('NO ACTIVO'),
	orden('PRIMARIO'),
	X = 'AMORFO'.

personalidad(X) :- 
    emotivo('NO EMOTIVO'),
	activo('NO ACTIVO'),
	orden('SECUNDARIO'),
	X = 'APATICO'.