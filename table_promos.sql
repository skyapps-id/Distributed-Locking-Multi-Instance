CREATE TABLE public.promos (
	uuid varchar NULL,
	code varchar NULL,
	qouta int8 NULL,
	updated_at timestamptz NULL
);

INSERT INTO public.promos
(uuid, code, qouta, updated_at)
VALUES('1677201d-bb68-45f3-b75e-e50df901595d', 'abc', 100, '2023-11-12 22:17:13.115');