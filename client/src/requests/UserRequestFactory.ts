/*
 * COPYRIGHT(c) 2023 MONTA
 *
 * This file is part of Monta.
 *
 * The Monta software is licensed under the Business Source License 1.1. You are granted the right
 * to copy, modify, and redistribute the software, but only for non-production use or with a total
 * of less than three server instances. Starting from the Change Date (November 16, 2026), the
 * software will be made available under version 2 or later of the GNU General Public License.
 * If you use the software in violation of this license, your rights under the license will be
 * terminated automatically. The software is provided "as is," and the Licensor disclaims all
 * warranties and conditions. If you use this license's text or the "Business Source License" name
 * and trademark, you must comply with the Licensor's covenants, which include specifying the
 * Change License as the GPL Version 2.0 or a compatible license, specifying an Additional Use
 * Grant, and not modifying the license in any other way.
 */

import { User } from "@/types/user";
import axios from "@/lib/axiosConfig";

/**
 * Return the user's job title information
 * @param user
 *
 * @returns {Promise<any>}
 */
export async function getUserJobTitle(user: User): Promise<any> {
  const response = await axios.get(`/job_titles/${user.profile?.job_title}/`);
  return response.data;
}

/***
 * Return the user's organization information
 * @param user
 *
 * @returns {Promise<any>}
 */
export async function getUserOrganization(user: User): Promise<any> {
  const response = await axios.get(`/organizations/${user.organization}/`);
  return response.data;
}

/**
 * Return the user's department information
 * @param user
 *
 * @returns {Promise<any>}
 */
export async function getUserDepartment(user: User): Promise<any> {
  const response = await axios.get(`/departments/${user?.department}/`);
  return response.data;
}

/**
 * Return the user's details
 * @param id
 *
 * @returns {Promise<any>}
 */
export async function getUserDetails(id: string): Promise<any> {
  const response = await axios.get(`/users/${id}/`);
  return response.data;
}

/**
 * Return a group of users.
 *
 * @returns {Promise<any>}
 */
export async function getUsers(): Promise<any> {
  const response = await axios.get("/users/");
  return response.data;
}
