/*
 * COPYRIGHT(c) 2024 Trenova
 *
 * This file is part of Trenova.
 *
 * The Trenova software is licensed under the Business Source License 1.1. You are granted the right
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

import { IChoiceProps } from "@/types";
import { TableOptionProps } from "@/types/tables";
import { CircleIcon, MinusCircledIcon } from "@radix-ui/react-icons";

/** Type for Account Type Choices */
export type AccountTypeChoiceProps =
  | "ASSET"
  | "LIABILITY"
  | "EQUITY"
  | "REVENUE"
  | "EXPENSE";

export const accountTypeChoices = [
  { value: "ASSET", label: "Asset" },
  { value: "LIABILITY", label: "Liability" },
  { value: "EQUITY", label: "Equity" },
  { value: "REVENUE", label: "Revenue" },
  { value: "EXPENSE", label: "Expense" },
] satisfies ReadonlyArray<IChoiceProps<AccountTypeChoiceProps>>;

export const tableAccountTypeChoices = [
  { value: "ASSET", label: "Asset" },
  { value: "LIABILITY", label: "Liability" },
  { value: "EQUITY", label: "Equity" },
  { value: "REVENUE", label: "Revenue" },
  { value: "EXPENSE", label: "Expense" },
] satisfies TableOptionProps[];

/** Type for Cash Flow Type Choices */
export type CashFlowTypeChoiceProps = "OPERATING" | "INVESTING" | "FINANCING";

export const cashFlowTypeChoices = [
  { value: "OPERATING", label: "Operating" },
  { value: "INVESTING", label: "Investing" },
  { value: "FINANCING", label: "Financing" },
] satisfies ReadonlyArray<IChoiceProps<CashFlowTypeChoiceProps>>;

/** Type for Account Sub Type Choices */
export type AccountSubTypeChoiceProps =
  | "CURRENT_ASSET"
  | "FIXED_ASSET"
  | "OTHER_ASSET"
  | "CURRENT_LIABILITY"
  | "LONG_TERM_LIABILITY"
  | "EQUITY"
  | "REVENUE"
  | "COST_OF_GOODS_SOLD"
  | "EXPENSE"
  | "OTHER_INCOME"
  | "OTHER_EXPENSE";

export const accountSubTypeChoices = [
  { value: "CURRENT_ASSET", label: "Current Asset" },
  { value: "FIXED_ASSET", label: "Fixed Asset" },
  { value: "OTHER_ASSET", label: "Other Asset" },
  { value: "CURRENT_LIABILITY", label: "Current Liability" },
  { value: "LONG_TERM_LIABILITY", label: "Long Term Liability" },
  { value: "EQUITY", label: "Equity" },
  { value: "REVENUE", label: "Revenue" },
  { value: "COST_OF_GOODS_SOLD", label: "Cost of Goods Sold" },
  { value: "EXPENSE", label: "Expense" },
  { value: "OTHER_INCOME", label: "Other Income" },
  { value: "OTHER_EXPENSE", label: "Other Expense" },
] satisfies ReadonlyArray<IChoiceProps<AccountSubTypeChoiceProps>>;

/** Type for Account Classification Choices */
export type AccountClassificationChoiceProps =
  | "BANK"
  | "CASH"
  | "ACCOUNTS_RECEIVABLE"
  | "ACCOUNTS_PAYABLE"
  | "INVENTORY"
  | "OTHER_CURRENT_ASSET"
  | "FIXED_ASSET";

export const accountClassificationChoices = [
  { value: "BANK", label: "Bank" },
  { value: "CASH", label: "Cash" },
  { value: "ACCOUNTS_RECEIVABLE", label: "Accounts Receivable" },
  { value: "ACCOUNTS_PAYABLE", label: "Accounts Payable" },
  { value: "INVENTORY", label: "Inventory" },
  { value: "OTHER_CURRENT_ASSET", label: "Other Current Asset" },
  { value: "FIXED_ASSET", label: "Fixed Asset" },
] satisfies ReadonlyArray<IChoiceProps<AccountClassificationChoiceProps>>;

/** Types for Job Function Choices */
export type JobFunctionChoiceProps =
  | "MANAGER"
  | "MANAGEMENT_TRAINEE"
  | "SUPERVISOR"
  | "DISPATCHER"
  | "BILLING"
  | "FINANCE"
  | "SAFETY"
  | "SYS_ADMIN"
  | "TEST";

export const jobFunctionChoices = [
  { value: "MANAGER", label: "Manager" },
  { value: "MANAGEMENT_TRAINEE", label: "Management Trainee" },
  { value: "SUPERVISOR", label: "Supervisor" },
  { value: "DISPATCHER", label: "Dispatcher" },
  { value: "BILLING", label: "Billing" },
  { value: "FINANCE", label: "Finance" },
  { value: "SAFETY", label: "Safety" },
  { value: "SYS_ADMIN", label: "System Administrator" },
  { value: "TEST", label: "Test Job Function" },
] satisfies ReadonlyArray<IChoiceProps<JobFunctionChoiceProps>>;

export const DayOfWeekChoices = [
  { value: 0, label: "Monday" },
  { value: 1, label: "Tuesday" },
  { value: 2, label: "Wednesday" },
  { value: 3, label: "Thursday" },
  { value: 4, label: "Friday" },
  { value: 5, label: "Saturday" },
  { value: 6, label: "Sunday" },
] satisfies IChoiceProps<number>[];

export type ServiceIncidentControlChoiceProps =
  | "Never"
  | "Pickup"
  | "Delivery"
  | "PickupAndDelivery"
  | "AllExceptShipper";

export enum ServiceIncidentControlEnum {
  Never = "Never",
  Pickup = "Pickup",
  Delivery = "Delivery",
  PickupAndDelivery = "PickupAndDelivery",
  AllExceptShipper = "AllExceptShipper",
}

export const serviceIncidentControlChoices = [
  { value: "Never", label: "Never" },
  { value: "Pickup", label: "Pickup" },
  { value: "Delivery", label: "Delivery" },
  { value: "PickupAndDelivery", label: "Pickup and Delivery" },
  { value: "AllExceptShipper", label: "All except shipper" },
] satisfies ReadonlyArray<IChoiceProps<ServiceIncidentControlChoiceProps>>;

/** Type for Date Format Choices */
export type DateFormatChoiceProps =
  | "InvoiceDateFormatMDY"
  | "InvoiceDateFormatDMY"
  | "InvoiceDateFormatYMD"
  | "InvoiceDateFormatYDM"
  | "";

export const dateFormatChoices = [
  { value: "InvoiceDateFormatMDY", label: "01/02/2006" },
  { value: "InvoiceDateFormatDMY", label: "02/01/2006" },
  { value: "InvoiceDateFormatYMD", label: "2006/02/01" },
  { value: "InvoiceDateFormatYDM", label: "2006/01/02" },
] satisfies ReadonlyArray<IChoiceProps<DateFormatChoiceProps>>;

/** Type for Route Avoidance Choices */
type RouteAvoidanceChoiceProps = "tolls" | "highways" | "ferries";

export const routeAvoidanceChoices = [
  { value: "tolls", label: "Tolls" },
  { value: "highways", label: "Highways" },
  { value: "ferries", label: "Ferries" },
] satisfies ReadonlyArray<IChoiceProps<RouteAvoidanceChoiceProps>>;

/** Type for Route Model Choices */
export type RouteModelChoiceProps = "best_guess" | "optimistic" | "pessimistic";

export const routeModelChoices = [
  { value: "best_guess", label: "Best Guess" },
  { value: "optimistic", label: "Optimistic" },
  { value: "pessimistic", label: "Pessimistic" },
] satisfies ReadonlyArray<IChoiceProps<RouteModelChoiceProps>>;

/** Type for Route Distance Unit Choices */
export type RouteDistanceUnitProps = "M" | "I";

export const routeDistanceUnitChoices = [
  { value: "M", label: "Metric" },
  { value: "I", label: "Imperial" },
] satisfies ReadonlyArray<IChoiceProps<RouteDistanceUnitProps>>;

/** Type for Distance Method Choices */
export type DistanceMethodChoiceProps = "G" | "T";

export const distanceMethodChoices = [
  { value: "G", label: "Google" },
  { value: "T", label: "Trenova" },
] satisfies ReadonlyArray<IChoiceProps<DistanceMethodChoiceProps>>;

/** Type for Email Protocol Choices */
export type EmailProtocolChoiceProps = "TLS" | "SSL" | "UNENCRYPTED";

export const emailProtocolChoices = [
  { value: "TLS", label: "TLS" },
  { value: "SSL", label: "SSL" },
  { value: "UNENCRYPTED", label: "Unencrypted" },
] satisfies ReadonlyArray<IChoiceProps<EmailProtocolChoiceProps>>;

/** Type for Feasibility Operator Choices */
export type FeasibilityOperatorChoiceProps =
  | "Eq"
  | "Ne"
  | "Gt"
  | "Gte"
  | "Lt"
  | "Lte";

export const feasibilityOperatorChoices = [
  { value: "Eq", label: "Equals" },
  { value: "Ne", label: "Not Equals" },
  { value: "Gt", label: "Greater Than" },
  { value: "Gte", label: "Greater Than or Equal To" },
  { value: "Lt", label: "Less Than" },
  { value: "Lte", label: "Less Than or Equal To" },
] satisfies ReadonlyArray<IChoiceProps<FeasibilityOperatorChoiceProps>>;

/** Type for Equipment Class Choices */
export type EquipmentClassChoiceProps =
  | "UNDEFINED"
  | "CAR"
  | "VAN"
  | "PICKUP"
  | "WALK_IN"
  | "STRAIGHT"
  | "TRACTOR"
  | "TRAILER";

export const equipmentClassChoices = [
  { value: "UNDEFINED", label: "Undefined" },
  { value: "CAR", label: "Car" },
  { value: "VAN", label: "Van" },
  { value: "PICKUP", label: "Pickup" },
  { value: "WALK_IN", label: "Walk In" },
  { value: "STRAIGHT", label: "Straight" },
  { value: "TRACTOR", label: "Tractor" },
  { value: "TRAILER", label: "Trailer" },
] satisfies ReadonlyArray<IChoiceProps<EquipmentClassChoiceProps>>;

/** Type for Unit of Measure Choices */
export type UnitOfMeasureChoiceProps =
  | "PALLET"
  | "TOTE"
  | "DRUM"
  | "CYLINDER"
  | "CASE"
  | "AMPULE"
  | "BAG"
  | "BOTTLE"
  | "PAIL"
  | "PIECES"
  | "ISO_TANK";

export const UnitOfMeasureChoices = [
  { value: "PALLET", label: "Pallet" },
  { value: "TOTE", label: "Tote" },
  { value: "DRUM", label: "Drum" },
  { value: "CYLINDER", label: "Cylinder" },
  { value: "CASE", label: "Case" },
  { value: "AMPULE", label: "Ampule" },
  { value: "BAG", label: "Bag" },
  { value: "BOTTLE", label: "Bottle" },
  { value: "PAIL", label: "Pail" },
  { value: "PIECES", label: "Pieces" },
  { value: "ISO_TANK", label: "ISO Tank" },
] satisfies ReadonlyArray<IChoiceProps<UnitOfMeasureChoiceProps>>;

/** Type for Hazardous Class Choices */
export type PackingGroupChoiceProps = "I" | "II" | "III";

export const packingGroupChoices = [
  { value: "I", label: "I" },
  { value: "II", label: "II" },
  { value: "III", label: "III" },
] satisfies ReadonlyArray<IChoiceProps<PackingGroupChoiceProps>>;

/** Type for Hazardous Class Choices */
export type HazardousClassChoiceProps =
  | "1.1"
  | "1.2"
  | "1.3"
  | "1.4"
  | "1.5"
  | "1.6"
  | "2.1"
  | "2.2"
  | "2.3"
  | "3"
  | "4.1"
  | "4.2"
  | "4.3"
  | "5.1"
  | "5.2"
  | "6.1"
  | "6.2"
  | "7"
  | "8"
  | "9";

export const hazardousClassChoices = [
  { value: "1.1", label: "Division 1.1: Mass Explosive Hazard" },
  { value: "1.2", label: "Division 1.2: Projection Hazard" },
  {
    value: "1.3",
    label: "Division 1.3: Fire and/or Minor Blast/Minor Projection Hazard",
  },
  { value: "1.4", label: "Division 1.4: Minor Explosion Hazard" },
  {
    value: "1.5",
    label: "Division 1.5: Very Insensitive With Mass Explosion Hazard",
  },
  {
    value: "1.6",
    label: "Division 1.6: Extremely Insensitive; No Mass Explosion Hazard",
  },
  { value: "2.1", label: "Division 2.1: Flammable Gases" },
  { value: "2.2", label: "Division 2.2: Non-Flammable Gases" },
  { value: "2.3", label: "Division 2.3: Poisonous Gases" },
  { value: "3", label: "Division 3: Flammable Liquids" },
  { value: "4.1", label: "Division 4.1: Flammable Solids" },
  { value: "4.2", label: "Division 4.2: Spontaneously Combustible Solids" },
  { value: "4.3", label: "Division 4.3: Dangerous When Wet" },
  { value: "5.1", label: "Division 5.1: Oxidizing Substances" },
  { value: "5.2", label: "Division 5.2: Organic Peroxides" },
  { value: "6.1", label: "Division 6.1: Toxic Substances" },
  { value: "6.2", label: "Division 6.2: Infectious Substances" },
  { value: "7", label: "Division 7: Radioactive Material" },
  { value: "8", label: "Division 8: Corrosive Substances" },
  {
    value: "9",
    label: "Division 9: Miscellaneous Hazardous Substances and Articles",
  },
] satisfies ReadonlyArray<IChoiceProps<HazardousClassChoiceProps>>;

/* Transaction Type Choice Type */
export type TransactionTypeChoiceType = "REVENUE" | "EXPENSE";
export const transactionTypeChoices = [
  { value: "REVENUE", label: "Revenue" },
  { value: "EXPENSE", label: "Expense" },
] satisfies ReadonlyArray<IChoiceProps<TransactionTypeChoiceType>>;

/* Transaction Status Choice Type */
export type TransactionStatusChoiceType =
  | "PENDING"
  | "COMPLETED"
  | "FAILED"
  | "PENDING_RECON";

export const transactionStatusChoices = [
  { value: "PENDING", label: "Pending" },
  { value: "PENDING_RECON", label: "Pending Reconciliation" },
  { value: "COMPLETED", label: "Completed" },
  { value: "FAILED", label: "Failed" },
] satisfies ReadonlyArray<IChoiceProps<TransactionStatusChoiceType>>;

/* Threshold Action Choice Type */
export type ThresholdActionChoiceType = "HALT" | "WARN";
export const thresholdActionChoices = [
  { value: "HALT", label: "Halt" },
  { value: "WARN", label: "Warn" },
] satisfies ReadonlyArray<IChoiceProps<ThresholdActionChoiceType>>;

/** Automatic Journal Entry Choices */
export type AutomaticJournalEntryChoiceType =
  | "OnShipmentBill"
  | "OnReceiptOfPayment"
  | "OnExpenseRecognition"
  | "";

export const automaticJournalEntryChoices = [
  {
    value: "OnShipmentBill",
    label: "Auto create entry when shipment is billed",
  },
  {
    value: "OnReceiptOfPayment",
    label: "Auto create entry on receipt of payment",
  },
  {
    value: "OnExpenseRecognition",
    label: "Auto create entry when an expense is recognized",
  },
] satisfies ReadonlyArray<IChoiceProps<AutomaticJournalEntryChoiceType>>;

/**
 * Returns status choices for a select input.
 */
type TStatusChoiceProps = "A" | "I";

/**
 * Returns status choices for a select input.
 * @returns An array of status choices.
 */
export const statusChoices = [
  { value: "A", label: "Active", color: "#15803d" },
  { value: "I", label: "Inactive", color: "#b91c1c" },
] satisfies ReadonlyArray<IChoiceProps<TStatusChoiceProps>>;

/** Entry methods for shipments */
export type EntryMethodChoiceProps = "MANUAL" | "EDI" | "API";

export const entryMethodChoices = [
  { value: "MANUAL", label: "Manual" },
  { value: "EDI", label: "EDI" },
  { value: "API", label: "API" },
] satisfies ReadonlyArray<IChoiceProps<EntryMethodChoiceProps>>;

/** Returns status choices for a select input. */
export type ShipmentStatusChoiceProps = "N" | "P" | "C" | "H" | "B" | "V";

/**
 * Returns shipment status choices for a select input.
 * @returns An array of shipment status choices.
 */
export const shipmentStatusChoices = [
  { value: "N", label: "New", color: "#16a34a" },
  { value: "P", label: "In Progress", color: "#ca8a04" },
  { value: "C", label: "Completed", color: "#9333ea" },
  { value: "H", label: "On Hold", color: "#2563eb" },
  { value: "B", label: "Billed", color: "#0891b2" },
  { value: "V", label: "Voided", color: "#dc2626" },
] satisfies ReadonlyArray<IChoiceProps<ShipmentStatusChoiceProps>>;

export type ShipmentEntryMethodChoices = "MANUAL" | "EDI" | "API";

export const shipmentSourceChoices = [
  { value: "MANUAL", label: "Manual" },
  { value: "EDI", label: "EDI" },
  { value: "API", label: "API" },
] satisfies ReadonlyArray<IChoiceProps<ShipmentEntryMethodChoices>>;

/**
 * Returns code type choices for a select input.
 */
export type CodeTypeProps = "VOIDED" | "CANCELLED";

/**
 * Returns code type choices for a select input.
 * @returns An array of code type choices.
 */
export const codeTypeChoices = [
  { value: "VOIDED", label: "Voided" },
  { value: "CANCELLED", label: "Cancelled" },
] satisfies ReadonlyArray<IChoiceProps<CodeTypeProps>>;

/**
 * Returns boolean yes & no choices for a select input.
 * @returns An array of yes & no choices.
 */
export const booleanStatusChoices: ReadonlyArray<IChoiceProps<boolean>> = [
  { value: true, label: "Active" },
  { value: false, label: "Inactive" },
];

/* Type for Database Actions */
export type DatabaseActionChoicesProps =
  | "INSERT"
  | "UPDATE"
  | "DELETE"
  | "BOTH";

export const databaseActionChoices = [
  { value: "INSERT", label: "Insert" },
  { value: "UPDATE", label: "Update" },
  { value: "DELETE", label: "Delete" },
  { value: "BOTH", label: "Both" },
] satisfies ReadonlyArray<IChoiceProps<DatabaseActionChoicesProps>>;

/* Type for Table Change Alert Source */
export type SourceChoicesProps = "KAFKA" | "POSTGRES";

export const sourceChoices = [
  { value: "KAFKA", label: "Kafka" },
  { value: "POSTGRES", label: "Postgres" },
] satisfies ReadonlyArray<IChoiceProps<SourceChoicesProps>>;

/** Type for RatingMethodW for Shipment */
export type RatingMethodChoiceProps = "F" | "PM" | "PS" | "PP" | "O";

export const ratingMethodChoices = [
  { value: "F", label: "Flat" },
  { value: "PM", label: "Per Mile" },
  { value: "PS", label: "Per Stop" },
  { value: "PP", label: "Per Pound" },
  { value: "O", label: "Other" },
] satisfies ReadonlyArray<IChoiceProps<RatingMethodChoiceProps>>;

/** Type for Shipment Stop Choices */
export type ShipmentStopChoices = "P" | "SP" | "SD" | "D" | "DO";

/** Returns shipment stop choices for a select input. */
export const shipmentStopChoices = [
  { value: "P", label: "Pickup" },
  { value: "SP", label: "Stop Pickup" },
  { value: "SD", label: "Stop Delivery" },
  { value: "D", label: "Delivery" },
  { value: "DO", label: "Drop Off" },
] satisfies ReadonlyArray<IChoiceProps<ShipmentStopChoices>>;

/** Type for Comment Type Severity */
export type SeverityChoiceProps = "H" | "M" | "L";

/** Returns comment type severity choices for a select input. */
export const severityChoices = [
  { value: "H", label: "High", color: "#dc2626" },
  { value: "M", label: "Medium", color: "#2563eb" },
  { value: "L", label: "Low", color: "#15803d" },
] satisfies ReadonlyArray<IChoiceProps<SeverityChoiceProps>>;

/**
 * Type for timezone choices
 */
export type TimezoneChoices =
  | "AmericaLosAngeles"
  | "AmericaDenver"
  | "AmericaChicago"
  | "AmericaNewYork";

/**
 * Returns timezone choices for a select input
 * @returns An array of timezone choices.
 */
export const timezoneChoices: ReadonlyArray<IChoiceProps<TimezoneChoices>> = [
  { value: "AmericaLosAngeles", label: "America/Los_Angeles" },
  { value: "AmericaDenver", label: "America/Denver" },
  { value: "AmericaChicago", label: "America/Chicago" },
  { value: "AmericaNewYork", label: "America/New_York" },
];

/**
 * Returns status choices when using TableFacetedFilters.
 * @returns An array of table faceted filter choices.
 */
export const tableStatusChoices = [
  {
    value: "A",
    label: "Active",
    icon: CircleIcon,
  },
  {
    value: "I",
    label: "Inactive",
    icon: MinusCircledIcon,
  },
] satisfies TableOptionProps[];

/**
 * Returns yes & no choices for a select input.
 * @returns An array of yes & no choices.
 */
export const yesAndNoChoices = [
  { value: "Y", label: "Yes" },
  { value: "N", label: "No" },
];

export type SegregationTypeChoiceProps = "O" | "X";
/**
 * Returns Segregation Type choices when using TableFacetedFilters.
 * @returns An array of table faceted filter choices.
 */
export const segregationTypeChoices = [
  {
    value: "O",
    label: "Allowed With Conditions",
    icon: CircleIcon,
  },
  {
    value: "X",
    label: "Not Allowed",
    icon: MinusCircledIcon,
  },
] satisfies TableOptionProps[];
