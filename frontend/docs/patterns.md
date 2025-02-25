# Patterns

This contains the patterns that we follow in the Fleet UI.

> NOTE: There are always exceptions to the rules, but we try as much as possible to
follow these patterns unless a specific use case calls for something else. These
should be discussed within the team and documented before merged.

## Table of contents
  - [Typing](#typing)
  - [Utilities](#utilities)
  - [Components](#components)
  - [React Hooks](#react-hooks)
  - [React Context](#react-context)
  - [Fleet API Calls](#fleet-api-calls)
  - [Page Routing](#page-routing)
  - [Styles](#styles)
  - [Other](#other)

## Typing
All Javascript and React files use Typescript, meaning the extensions are `.ts` and `.tsx`. Here are the guidelines on how we type at Fleet:

- Use *[global entity interfaces](../README.md#interfaces)* when interfaces are used multiple times across the app
- Use *local interfaces* when typing entities limited to the specific page or component
- Local interfaces for page and component props

  ```typescript
  // page
  interface IPageProps {
    prop1: string;
    prop2: number;
    ...
  }

  // Note: Destructure props in page/component signature
  const PageOrComponent = ({
    prop1,
    prop2,
  }: IPageProps) => {

    return (
      // ...
    );
  };
  ```

- Local states
```typescript
const [item, setItem] = useState("");
```

- Fetch function signatures (i.e. `react-query`)
```typescript
useQuery<IHostResponse, Error, IHost>(params)
```

- Custom functions, including callbacks
```typescript
const functionWithTableName = (tableName: string): boolean => {
  // do something
};
```

## Utilities

### Named exports

We export individual utility functions and avoid exporting default objects when exporting utilities.

```ts

// good
export const replaceNewLines = () => {...}

// bad
export default {
  replaceNewLines
}
```

## Components

### React Functional Components

We use functional components with React instead of class comonents. We do this
as this allows us to use hooks to better share common logic between components.

### Page Component Pattern

When creating a **top level page** (e.g. dashboard page, hosts page, policies page)
we wrap that page's content inside components `MainContent` and
`SidePanelContent` if a sidebar is needed.

These components encapsulate the styling used for laying out content and also
handle rendering of common UI shared across all pages (current this is only the
sandbox expiry message with more to come).

```typescript
/** An example of a top level page utilising MainConent and SidePanel content */
const PackComposerPage = ({ router }: IPackComposerPageProps): JSX.Element => {
  // ...

  return (
    <>
      <MainContent className={baseClass}>
        <PackForm
          className={`${baseClass}__pack-form`}
          handleSubmit={handleSubmit}
          onFetchTargets={onFetchTargets}
          selectedTargetsCount={selectedTargetsCount}
          isPremiumTier={isPremiumTier}
        />
      </MainContent>
      <SidePanelContent>
        <PackInfoSidePanel />
      </SidePanelContent>
    </>
  );
};

export default PackComposerPage;
```

## React Hooks

[Hooks](https://reactjs.org/docs/hooks-intro.html) are used to track state and use other features
of React. Hooks are only allowed in functional components, which are created like so:

```typescript
import React, { useState, useEffect } from "React";

const PageOrComponent = (props) => {
  const [item, setItem] = useState("");

  // runs only on first mount (replaces componentDidMount)
  useEffect(() => {
    // do something
  }, []);

  // runs only when `item` changes (replaces componentDidUpdate)
  useEffect(() => {
    // do something
  }, [item]);

  return (
    // ...
  );
};
```

> NOTE: Other hooks are available per [React's documentation](https://reactjs.org/docs/hooks-intro.html).

## React context

[React context](https://reactjs.org/docs/context.html) is a state management store. It stores
data that is desired and allows for retrieval of that data in whatever component is in need.
View currently working contexts in the [context directory](../context).

## Fleet API calls

The [services](../services) directory stores all API calls and is to be used in two ways:
- A direct `async/await` assignment
- Using `react-query` if requirements call for loading data right away or based on dependencies.

Examples below:

**Direct assignment**

```typescript
// page
import ...
import queriesAPI from "services/entities/queries";

const PageOrComponent = (props) => {
  const doSomething = async () => {
    try {
      const response = await queriesAPI.load(param);
      // do something
    } catch(error) {
      console.error(error);
      // maybe trigger renderFlash
    }
  };

  return (
    // ...
  );
};
```

**React Query**

[react-query](https://react-query.tanstack.com/overview) is a data-fetching library that
gives us the ability to fetch, cache, sync and update data with a myriad of options and properties.

```typescript
import ...
import { useQuery, useMutation } from "react-query";
import queriesAPI from "services/entities/queries";

const PageOrComponent = (props) => {
  // retrieve the query based on page/component load
  // and dependencies for when to refetch
  const {
    isLoading,
    data,
    error,
    ...otherProps,
  } = useQuery<IResponse, Error, IData>(
    "query",
    () => queriesAPI.load(param),
    {
      ...options
    }
  );

  // `props` is a bucket of properties that can be used when
  // updating data. for example, if you need to know whether
  // a mutation is loading, there is a prop for that.
  const { ...props } = useMutation((formData: IForm) =>
    queriesAPI.create(formData)
  );

  return (
    // ...
  );
};
```

## Page routing

We use React Router directly to navigate between pages. For page components,
React Router (v3) supplies a `router` prop that can be easily accessed.
When needed, the `router` object contains a `push` function that redirects
a user to whatever page desired. For example:

```typescript
// page
import PATHS from "router/paths";
import { InjectedRouter } from "react-router/lib/Router";

interface IPageProps {
  router: InjectedRouter; // v3
}

const PageOrComponent = ({
  router,
}: IPageProps) => {
  const doSomething = () => {
    router.push(PATHS.HOME);
  };

  return (
    // ...
  );
};
```

## Styles

Below are a few need-to-knows about what's available in Fleet's CSS:

### Modals

1) When creating a modal with a form inside, the action buttons (cancel, save, delete, etc.) should
   be wrapped in the `modal-cta-wrap` class to keep unified styles.

### Forms

1) When creating a form, **not** in a modal, use the class `${baseClass}__button-wrap` for the
   action buttons (cancel, save, delete, etc.) and proceed to style as needed.

## Other

### Local states

Our first line of defense for state management is local states (i.e. `useState`). We
use local states to keep pages/components separate from one another and easy to
maintain. If states need to be passed to direct children, then prop-drilling should
suffice as long as we do not go more than two levels deep. Otherwise, if states need
to be used across multiple unrelated components or 3+ levels from a parent,
then the [app's context](#react-context) should be used.

### File size

The recommend line limit per page/component is 500 lines. This is only a recommendation.
Larger files are to be split into multiple files if possible.
