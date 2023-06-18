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
import {
  Box,
  createStyles,
  Modal,
  rem,
  SimpleGrid,
  Skeleton,
  Stack,
  Text,
} from "@mantine/core";
import { useQuery, useQueryClient } from "react-query";
import { User } from "@/types/apps/accounts";
import {
  getDepartmentDetails,
  getOrganizationDetails,
} from "@/requests/OrganizationRequestFactory";
import React from "react";

interface ViewUserModalProps {
  onClose: () => void;
  opened: boolean;
  user: User | null;
}

const useStyles = createStyles((theme) => {
  const BREAKPOINT = theme.fn.smallerThan("sm");

  return {
    fields: {
      marginTop: rem(10),
    },
    control: {
      [BREAKPOINT]: {
        flex: 1,
      },
    },
    text: {
      color: theme.colorScheme === "dark" ? "white" : "black",
    },
    div: {
      marginBottom: rem(10),
    },
  };
});

export const ViewUserModal: React.FC<ViewUserModalProps> = ({
  onClose,
  opened,
  user,
}) => {
  const { classes } = useStyles();
  const queryClient = useQueryClient();

  const { data: organizationData, isLoading: isOrganizationDataLoading } =
    useQuery({
      queryKey: ["organization", user?.organization],
      queryFn: () => {
        if (!user || !user.organization) {
          return Promise.resolve(null);
        }
        return getOrganizationDetails(user.organization);
      },
      enabled: opened && !!user?.organization,
      initialData: () => {
        return queryClient.getQueryData(["organization", user?.id]);
      },
      staleTime: Infinity, // never refetch
    });

  const { data: departmentData, isLoading: isDepartmentDataLoading } = useQuery(
    {
      queryKey: ["department", user?.department],
      queryFn: () => {
        if (!user || !user.department) {
          return Promise.resolve(null);
        }
        return getDepartmentDetails(user.department);
      },
      enabled: opened && !!user?.department,
      initialData: () => {
        return queryClient.getQueryData(["department", user?.department]);
      },
      staleTime: Infinity, // never refetch
    }
  );

  const isUserDataLoading =
    isOrganizationDataLoading || isDepartmentDataLoading;

  return (
    <>
      <Modal.Root opened={opened} onClose={onClose} size="lg" centered>
        <Modal.Overlay />
        <Modal.Content>
          <Modal.Header>
            <Modal.Title>View User</Modal.Title>
            <Modal.CloseButton />
          </Modal.Header>
          <Modal.Body>
            {isUserDataLoading ? (
              <Stack>
                <Skeleton height={300} />
              </Stack>
            ) : (
              <>
                <SimpleGrid cols={3} spacing="sm" verticalSpacing="sm">
                  <Box>
                    <Text className={classes.text}>Organization</Text>
                    <Text color="dimmed" fz="sm">
                      {organizationData?.name}
                    </Text>
                  </Box>
                  <Box>
                    <Text className={classes.text}>Username</Text>
                    <Text color="dimmed" fz="sm">
                      {user?.username}
                    </Text>
                  </Box>
                  <Box>
                    <Text className={classes.text}>Email</Text>
                    <Text color="dimmed" fz="sm">
                      {user?.email}
                    </Text>
                  </Box>
                </SimpleGrid>
                <SimpleGrid cols={3} spacing="sm" verticalSpacing="sm" mt={15}>
                  <Box>
                    <Text className={classes.text}>Department</Text>
                    <Text color="dimmed" fz="sm">
                      {departmentData?.name}
                    </Text>
                  </Box>
                  <Box>
                    <Text className={classes.text}>Username</Text>
                    <Text color="dimmed" fz="sm">
                      {user?.username}
                    </Text>
                  </Box>
                  <Box>
                    <Text className={classes.text}>Email</Text>
                    <Text color="dimmed" fz="sm">
                      {user?.email}
                    </Text>
                  </Box>
                </SimpleGrid>
              </>
            )}
          </Modal.Body>
        </Modal.Content>
      </Modal.Root>
    </>
  );
};
