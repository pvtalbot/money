<script setup>
// Pinia
import { useUserStore } from '@/stores/user.js';
import { useDrawerStore } from '@/stores/drawer.js';

// External libraries
import dayjs from 'dayjs';
import utc from 'dayjs/plugin/utc';

// Vue
import { computed, ref } from 'vue';
import TransferList from '@/components/transfers/TransferList.vue';
import VueButton from '@/components/utils/VueButton.vue';
import CreateTransferForm from '@/components/transfers/CreateTransferForm.vue';
import DatePicker from '@/components/utils/DatePicker.vue';

dayjs.extend(utc);
const userStore = useUserStore();
const drawerStore = useDrawerStore();

const mainDatePicker = ref(null)
const mode = ref("")

const firstName = computed(() => userStore.user.firstName)
const COMPONENT_TO_DRAWER = "CreateTransferForm";

const initialDate = dayjs().utc().date(1).hour(0).minute(0).second(0).millisecond(0);
const dateToProp = computed(() => mainDatePicker == null || mainDatePicker.value == null ? initialDate : mainDatePicker.value.date);

const openCreateForm = v => {
  if (!['expense', 'revenue'].includes(v)) return;
  mode.value = v;
  drawerStore.registerComponent(COMPONENT_TO_DRAWER);
}
</script>

<template>
  <h1>Hello {{firstName}}, this is Expenses manager!</h1>
  <div class="transfers">
    <div class="datepicker">
      <DatePicker :initialDate="initialDate" ref="mainDatePicker"/>
    </div>
    <Teleport to="#teleport-component-to-drawer">
      <Transition name="component">
        <CreateTransferForm v-if="drawerStore.isCurrentComponentDisplayed(COMPONENT_TO_DRAWER)" :mode="mode"/>
      </Transition>
    </Teleport>
    <div class="revenues">
      <div class="create-revenue">
        <VueButton message="Add a revenue" @click="openCreateForm('revenue')"/>
      </div>
      <TransferList :initialDate="dateToProp" :mode="'revenue'"/>
    </div>
    <div class="expenses">
      <div class="create-expense" @click="openCreateForm('expense')">
        <VueButton message="Add an expense" />
      </div>
      <TransferList :initialDate="dateToProp" :mode="'expense'"/>
    </div>
  </div>
</template>

<style scoped>
.transfers {
  display: grid;
  grid-template-columns: 1fr 1fr;
  grid-template-rows: 50px 1fr;
  grid-template-areas: "header header" "left right";
}

.revenues, .expenses {
  margin: 10px auto auto auto;
}

.datepicker {
  grid-area: header;
  margin: auto;
}

.revenues {
  grid-area: left;
}

.expenses {
  grid-area: right;
}

.create-expense, .create-revenue {
  margin: auto;
  width: fit-content;
}

.component-leave-active, .component-enter-active {
  transition: all 0.5s ease;
}

.component-leave-to, .component-enter-from {
  transform: translateX(320px);
}
</style>