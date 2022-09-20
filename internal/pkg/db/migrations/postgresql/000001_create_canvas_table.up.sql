CREATE TABLE public.canvases (
    id uuid NOT NULL,
    name character varying(255) NOT NULL,
    contents jsonb NULL
);

ALTER TABLE ONLY public.canvases
    ADD CONSTRAINT canvas_pkey PRIMARY KEY (id);

CREATE UNIQUE INDEX idx_canvas_name ON canvases(name);
