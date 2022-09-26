// Apollo
import { useLazyQuery } from '@vue/apollo-composable';
import Me from '@/graphql/queries/CurrentUser.gql';
import GetAllErrors from '@/graphql/queries/GetAllErrors.gql';
import ValidateAuthToken from '@/graphql/queries/ValidateAuthToken.gql';

// Pinia
import { useExpenseStore } from '@/stores/expense';
import { useErrorStore } from '@/stores/errors';
import { useUserStore } from '@/stores/user';

// Vue
import { onMounted } from 'vue';

const AUTH_TOKEN = 'authToken';
const REFRESH_TOKEN = 'refreshToken';

const useCheckLocalStorageForToken = () =>
  new Promise((resolve, reject) => {
    // Apollo Query to check the validity of a token. Used on mounted
    const { result: tokenValidity, load, onResult: onTokenValidated } = useLazyQuery(ValidateAuthToken);

    // Checks if there is a token in local storage. If yes, checks validity.
    // If the token is still valid, logs the user in
    onMounted(() => {
      const authToken = getUserToken();
      if (authToken == null) reject();
      else {
        load(undefined, { authToken });
      }
    })

    onTokenValidated(() => {
      if (!tokenValidity.value.validateAuthToken) {
        removeUserToken();
        reject();
      } else {
        resolve();
      }
    })
  });

const useLoadCurrentUser = () => {
  const expenseStore = useExpenseStore();
  const userStore = useUserStore();

  const { result: currentUser, load, onResult } = useLazyQuery(Me);

  onResult(() => {
    if (!currentUser.value) return;
    userStore.$patch(store => {
      store.user.firstName = currentUser.value.me.firstName;
      store.user.lastName = currentUser.value.me.lastName;
      store.user.id = currentUser.value.me.id;
    })

    expenseStore.cacheExpensesCategories(currentUser.value.me.expensesCategories);
  })

  return load;
}

const useGetAllErrors = () => {
  const errorStore = useErrorStore();

  const { result: errorsList, load, onResult: onGetAllErrorsSuccess } = useLazyQuery(GetAllErrors);

  onGetAllErrorsSuccess(() => {
    errorStore.cacheErrors(errorsList.value.getAllErrors);
  })

  return load;
};

const storeUserToken = (token) => {
  if (typeof localStorage !== undefined && token) {
    localStorage.setItem(AUTH_TOKEN, token.authToken);
    localStorage.setItem(REFRESH_TOKEN, token.refreshToken);
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
  storeUserToken,
  useCheckLocalStorageForToken,
  useLoadCurrentUser,
};