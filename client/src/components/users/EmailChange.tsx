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

import React, { useState } from "react";
import { Text, Col, Button, Grid, SimpleGrid, Divider } from "@mantine/core";
import * as Yup from "yup";
import { useForm, yupResolver } from "@mantine/form";
import { notifications } from "@mantine/notifications";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faCheck } from "@fortawesome/pro-solid-svg-icons";
import { useQueryClient } from "react-query";
import axios from "@/lib/AxiosConfig";
import { ValidatedTextInput } from "../ui/fields/TextInput";
import { User } from "@/types/apps/accounts";
import { useFormStyles } from "@/styles/FormStyles";

interface EmailChangeFormProps {
  user: User;
}

export function EmailChangeForm({ user }: EmailChangeFormProps) {
  const [isEditing, setIsEditing] = useState<boolean>(false);
  const [loading, setLoading] = useState<boolean>(false);
  const { classes } = useFormStyles();
  const queryClient = useQueryClient();

  const startEditing = () => setIsEditing(true);
  const stopEditing = () => {
    form.reset();
    setIsEditing(false);
  };

  interface FormValues {
    email: string;
    currentPassword: string;
  }

  const schema = Yup.object().shape({
    email: Yup.string()
      .email("Invalid email address")
      .required("Email is required"),
    currentPassword: Yup.string().required("Password is required"),
  });

  const form = useForm<FormValues>({
    validate: yupResolver(schema),
    initialValues: {
      email: user.email,
      currentPassword: "",
    },
  });

  const submitForm = async (values: FormValues) => {
    setLoading(true);
    try {
      const response = await axios.post("change_email/", values);

      if (response.status === 200) {
        queryClient.invalidateQueries("users", { exact: true }).then(() => {
          stopEditing();
          notifications.show({
            title: "Email Changed",
            message: "Your email has been changed successfully",
            color: "green",
            withCloseButton: true,
            icon: <FontAwesomeIcon icon={faCheck} />,
          });
        });
      }
    } catch (error: any) {
      if (error.response) {
        const { data } = error.response;
        Object.keys(data).forEach((field) => {
          const message = data[field].join(" ");
          form.setFieldError(field, message);
        });
      }
    } finally {
      setLoading(false);
    }
  };

  return (
    <>
      <div className={classes.div}>
        {!isEditing ? (
          <div
            style={{
              display: "flex",
              alignItems: "center",
              flexWrap: "wrap",
            }}
          >
            <div>
              <Text size="sm" className={classes.text} weight={700}>
                Email Address
              </Text>
              <Text color="dimmed">{user.email}</Text>
            </div>
            <div
              style={{
                marginLeft: "auto",
              }}
            >
              <Button color="gray" variant="light" onClick={startEditing}>
                Change Email
              </Button>
            </div>
          </div>
        ) : (
          <Grid>
            <Col w="auto">
              <form onSubmit={form.onSubmit((values) => submitForm(values))}>
                <SimpleGrid cols={2} mb={20}>
                  <ValidatedTextInput
                    label="Enter New Email Address"
                    type="email"
                    variant="filled"
                    name="email"
                    form={form}
                    withAsterisk
                  />
                  <ValidatedTextInput
                    label="Enter Password"
                    type="password"
                    variant="filled"
                    name="currentPassword"
                    form={form}
                    withAsterisk
                  />
                </SimpleGrid>
                <Button type="submit" color="blue" mx="xs" loading={loading}>
                  Update Email
                </Button>
                <Button
                  type="button"
                  color="gray"
                  variant="light"
                  onClick={stopEditing}
                >
                  Cancel
                </Button>
              </form>
            </Col>
          </Grid>
        )}
      </div>
      <Divider my="sm" variant="dashed" />
    </>
  );
}
