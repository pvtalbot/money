import { ApolloClient, createHttpLink, InMemoryCache } from '@apollo/client/core';
import { onError } from '@apollo/client/link/error';
import { setContext } from '@apollo/client/link/context';
import { logErrorMessages } from '@vue/apollo-util';
import openEndpoints from './openEndpoints';

const AUTH_TOKEN = 'accessToken';

const cache = new InMemoryCache()

const httpLink = createHttpLink({
  uri: 'http://localhost/query',
})

// Some operations do not require to be logged in (login, validate access token).
// To ensure we don't get any errors, we never include the Auth header for those
const needAuth = (operation) => !!operation && !openEndpoints.includes(operation)

const authLink = setContext(({operationName}, {headers}) => {
  const token = localStorage.getItem(AUTH_TOKEN);

  if (needAuth(operationName) && !!token) {
    return {
      headers: {
        ...headers,
        Authorization: token,
      }
    }
  }

  return headers;
})

const errorLink = onError(error => {
  if (import.meta.env.MODE !== 'production') {
    logErrorMessages(error);
  }
})


export const createApolloClient = () => {
  return new ApolloClient({
    link: errorLink.concat(authLink.concat(httpLink)),
    cache,
    defaultOptions: {
      watchQuery: {
        fetchPolicy: 'no-cache',
      },
      query: {
        fetchPolicy: 'no-cache',
      },
      mutate: {
        fetchPolicy: 'no-cache',
      }
    }
  })
}
