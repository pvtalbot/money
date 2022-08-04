import { ApolloClient, createHttpLink, HttpLink, InMemoryCache } from '@apollo/client/core';
import { setContext } from '@apollo/client/link/context';

const AUTH_TOKEN = 'accessToken';

const cache = new InMemoryCache()

const httpLink = createHttpLink({
  uri: 'http://localhost/query',
})

const authLink = setContext((_, { headers }) => {
  const token = localStorage.getItem(AUTH_TOKEN);

  if (!!token) {
    return {
      headers: {
        ...headers,
        Authorization: token,
      }
    }
  }

  return headers;
})

export const createApolloClient = () => {
  return new ApolloClient({
    link: authLink.concat(httpLink),
    cache,
    defaultOptions: {
      fetchPolicy: 'no-cache',
    }
  })
}