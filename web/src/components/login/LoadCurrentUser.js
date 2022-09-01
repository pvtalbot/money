// Apollo
import { useLazyQuery, useQuery } from '@vue/apollo-composable';
import Me from '@/graphql/queries/CurrentUser.gql';
import GetAllErrors from '@/graphql/queries/GetAllErrors.gql';
import ValidateAccessToken from '@/graphql/queries/ValidateAccessToken.gql';

// Pinia
import { useExpenseStore } from '@/stores/expense';
import { useErrorStore } from '@/stores/errors';
import { useUserStore } from '@/stores/user';

// Vue
import { onMounted } from 'vue';

const AUTH_TOKEN = 'accessToken';

const useCheckLocalStorageForToken = () =>
  new Promise((resolve, reject) => {
    // Apollo Query to check the validity of a token. Used on mounted
    const { result: tokenValidity, load, onResult: onTokenValidated } = useLazyQuery(ValidateAccessToken);

    // Checks if there is a token in local storage. If yes, checks validity.
    // If the token is still valid, logs the user in
    onMounted(() => {
      const accessToken = getUserToken();
      if (accessToken == null) reject();
      else {
        load(undefined, { accessToken: accessToken });
      }
    })

    onTokenValidated(() => {
      if (!tokenValidity.value.validateAccessToken) {
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
  storeUserToken,
  useCheckLocalStorageForToken,
  useLoadCurrentUser,
};