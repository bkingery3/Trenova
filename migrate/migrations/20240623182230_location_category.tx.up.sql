-- COPYRIGHT(c) 2024 Trenova
--
-- This file is part of Trenova.
--
-- The Trenova software is licensed under the Business Source License 1.1. You are granted the right
-- to copy, modify, and redistribute the software, but only for non-production use or with a total
-- of less than three server instances. Starting from the Change Date (November 16, 2026), the
-- software will be made available under version 2 or later of the GNU General Public License.
-- If you use the software in violation of this license, your rights under the license will be
-- terminated automatically. The software is provided "as is," and the Licensor disclaims all
-- warranties and conditions. If you use this license's text or the "Business Source License" name
-- and trademark, you must comply with the Licensor's covenants, which include specifying the
-- Change License as the GPL Version 2.0 or a compatible license, specifying an Additional Use
-- Grant, and not modifying the license in any other way.

CREATE TABLE
    IF NOT EXISTS "location_categories" (
        "id" uuid NOT NULL DEFAULT uuid_generate_v4 (),
        "business_unit_id" uuid NOT NULL,
        "organization_id" uuid NOT NULL,
        "name" VARCHAR(50) NOT NULL,
        "description" TEXT,
        "color" VARCHAR(10),
        "version" BIGINT NOT NULL,
        "created_at" TIMESTAMPTZ NOT NULL DEFAULT current_timestamp,
        "updated_at" TIMESTAMPTZ NOT NULL DEFAULT current_timestamp,
        PRIMARY KEY ("id"),
        FOREIGN KEY ("organization_id") REFERENCES organizations ("id") ON UPDATE NO ACTION ON DELETE CASCADE,
        FOREIGN KEY ("business_unit_id") REFERENCES business_units ("id") ON UPDATE NO ACTION ON DELETE CASCADE
    );

--bun:split
CREATE UNIQUE INDEX IF NOT EXISTS "location_categories_name_organization_id_unq" ON "location_categories" (LOWER("name"), organization_id);

CREATE INDEX idx_location_categories_name ON location_categories (name);

CREATE INDEX idx_location_categories_org_bu ON location_categories (organization_id, business_unit_id);

CREATE INDEX idx_location_categories_description ON location_categories USING GIN (description gin_trgm_ops);

CREATE INDEX idx_location_categories_created_at ON location_categories (created_at);

-- bun:split
COMMENT ON COLUMN location_categories.id IS 'Unique identifier for the location category, generated as a UUID';

COMMENT ON COLUMN location_categories.business_unit_id IS 'Foreign key referencing the business unit that this location category belongs to';

COMMENT ON COLUMN location_categories.organization_id IS 'Foreign key referencing the organization that this location category belongs to';

COMMENT ON COLUMN location_categories.name IS 'The name of the location category';

COMMENT ON COLUMN location_categories.description IS 'A detailed description of the location category';

COMMENT ON COLUMN location_categories.color IS 'The color associated with the location category, represented as a string limited to 10 characters';

COMMENT ON COLUMN location_categories.created_at IS 'Timestamp of when the location category was created, defaults to the current timestamp';

COMMENT ON COLUMN location_categories.updated_at IS 'Timestamp of the last update to the location category, defaults to the current timestamp';