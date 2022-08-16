// Apollo
import { useLazyQuery, useQuery } from '@vue/apollo-composable';
import Me from '@/graphql/queries/CurrentUser.gql';
import GetAllErrors from '@/graphql/queries/GetAllErrors.gql';

// Pinia
import { useExpenseStore } from '@/stores/expense';
import { useErrorStore } from '@/stores/errors';
import { useUserStore } from '@/stores/user';

// Vue
import { useRouter } from 'vue-router';

const AUTH_TOKEN = 'accessToken';

const useLoadCurrentUser = () => {
  const expenseStore = useExpenseStore();
  const userStore = useUserStore();
  const router = useRouter();

  const { result: currentUser, load: loadCurrentUser, onResult: onCurrentUserSucceeded } = useLazyQuery(Me);

  onCurrentUserSucceeded(() => {
    if (!currentUser.value) return;
    userStore.$patch(store => {
      store.user.firstName = currentUser.value.me.firstName;
      store.user.lastName = currentUser.value.me.lastName;
    })

    expenseStore.cacheExpensesCategories(currentUser.value.me.expensesCategories);

    router.push({ name: 'home' });
  })

  return loadCurrentUser;
}

const useGetAllErrors = () => {
  const errorStore = useErrorStore();

  const { result: errorsList, onResult: onGetAllErrorsSuccess } = useQuery(GetAllErrors);

  onGetAllErrorsSuccess(() => {
    errorStore.cacheErrors(errorsList.value.getAllErrors);
  })
}

const storeUserToken = (token) => {
  if (typeof localStorage !== undefined && token) {
    localStorage.setItem(AUTH_TOKEN, token);
  }
}

const getUserToken = () => {
  if (!localStorage) return null;
  return localStorage.getItem(AUTH_TOKEN);
}

const removeUserToken = () => {
  if (!localStorage) return;
  localStorage.removeItem(AUTH_TOKEN);
}

export { 
  useGetAllErrors,
  useLoadCurrentUser,
  storeUserToken,
  getUserToken,
  removeUserToken,
};