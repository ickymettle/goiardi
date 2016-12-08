/*
 * Copyright (c) 2013-2016, Jeremy Bingham (<jeremy@goiardi.gl>)
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package role

/* MySQL funcs for roles */

import (
	"github.com/ickymettle/goiardi/datastore"
)

func (r *Role) saveMySQL() error {
	rlb, rlerr := datastore.EncodeBlob(&r.RunList)
	if rlerr != nil {
		return rlerr
	}
	erb, ererr := datastore.EncodeBlob(&r.EnvRunLists)
	if ererr != nil {
		return ererr
	}
	dab, daerr := datastore.EncodeBlob(&r.Default)
	if daerr != nil {
		return daerr
	}
	oab, oaerr := datastore.EncodeBlob(&r.Override)
	if oaerr != nil {
		return oaerr
	}
	tx, err := datastore.Dbh.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec("INSERT INTO roles (name, description, run_list, env_run_lists, default_attr, override_attr, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, NOW(), NOW()) ON DUPLICATE KEY UPDATE description = ?, run_list = ?, env_run_lists = ?, default_attr = ?, override_attr = ?, updated_at = NOW()", r.Name, r.Description, rlb, erb, dab, oab, r.Description, rlb, erb, dab, oab)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
