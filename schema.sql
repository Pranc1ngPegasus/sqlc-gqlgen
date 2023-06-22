CREATE TABLE organizations (
  id uuid NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE NOT NULL,
  updated_at TIMESTAMP WITH TIME ZONE NOT NULL,
  name character varying NOT NULL,
  PRIMARY KEY(id)
);

CREATE TABLE users (
  id uuid NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE NOT NULL,
  updated_at TIMESTAMP WITH TIME ZONE NOT NULL,
  organization_id uuid NOT NULL,
  name character varying NOT NULL,
  PRIMARY KEY(id)
);
