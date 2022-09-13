import { useEffect, useState, Fragment } from 'react'
import Link from 'next/link'
import Button from '@mui/material/Button'
import { AxiosError } from 'axios'
import { styled, useTheme, Theme, CSSObject } from '@mui/material/styles'
import MuiAppBar, { AppBarProps as MuiAppBarProps } from '@mui/material/AppBar'
import Toolbar from '@mui/material/Toolbar'
import Typography from '@mui/material/Typography'
import IconButton from '@mui/material/IconButton'
import MenuIcon from '@mui/icons-material/Menu'
import MuiDrawer from '@mui/material/Drawer'
import Divider from '@mui/material/Divider'
import List from '@mui/material/List'
import ChevronLeftIcon from '@mui/icons-material/ChevronLeft'
import ChevronRightIcon from '@mui/icons-material/ChevronRight'
// @ts-ignore
import { ToggleDarkMode } from '@shortlink-org/ui-kit'

import SearchForm from '../SearchForm'
import Notification from './notification'
import Profile from './profile'
// @ts-ignore
import secondMenu from './secondMenu'
import { mainListItems, secondaryListItems, adminListItems } from './listItems'
import ory from '../../pkg/sdk'

const drawerWidth = 240

interface AppBarProps extends MuiAppBarProps {
  open?: boolean
}

/* eslint-disable */
const Drawer = styled(MuiDrawer, {
  shouldForwardProp: (prop) => prop !== 'open',
})(({ theme, open }) => ({
  width: drawerWidth,
  flexShrink: 0,
  whiteSpace: 'nowrap',
  boxSizing: 'border-box',
  ...(open && {
    ...openedMixin(theme),
    '& .MuiDrawer-paper': openedMixin(theme),
  }),
  ...(!open && {
    ...closedMixin(theme),
    '& .MuiDrawer-paper': closedMixin(theme),
  }),
}))

const DrawerHeader = styled('div')(({ theme }) => ({
  display: 'flex',
  alignItems: 'center',
  justifyContent: 'flex-end',
  padding: theme.spacing(0, 1),
  // necessary for content to be below app bar
  ...theme.mixins.toolbar,
}))

const openedMixin = (theme: Theme): CSSObject => ({
  width: drawerWidth,
  transition: theme.transitions.create('width', {
    easing: theme.transitions.easing.sharp,
    duration: theme.transitions.duration.enteringScreen,
  }),
  overflowX: 'hidden',
})

const closedMixin = (theme: Theme): CSSObject => ({
  transition: theme.transitions.create('width', {
    easing: theme.transitions.easing.sharp,
    duration: theme.transitions.duration.leavingScreen,
  }),
  overflowX: 'hidden',
  width: `calc(${theme.spacing(7)} + 1px)`,
  [theme.breakpoints.up('sm')]: {
    width: `calc(${theme.spacing(8)} + 1px)`,
  },
})

const AppBar = styled(MuiAppBar, {
  shouldForwardProp: (prop) => prop !== 'open',
})<AppBarProps>(({ theme, open }) => ({
  zIndex: theme.zIndex.drawer + 1,
  transition: theme.transitions.create(['width', 'margin'], {
    easing: theme.transitions.easing.sharp,
    duration: theme.transitions.duration.leavingScreen,
  }),
  ...(open && {
    marginLeft: drawerWidth,
    width: `calc(100% - ${drawerWidth}px)`,
    transition: theme.transitions.create(['width', 'margin'], {
      easing: theme.transitions.easing.sharp,
      duration: theme.transitions.duration.enteringScreen,
    }),
  }),
}))

const Header = () => {
  const [session, setSession] = useState<string>(
    'No valid Ory Session was found.\nPlease sign in to receive one.',
  )
  const [hasSession, setHasSession] = useState<boolean>(false)

  useEffect(() => {
    ory
      .toSession()
      .then(({ data }) => {
        setSession(JSON.stringify(data, null, 2))
        setHasSession(true)
      })
      .catch((err: AxiosError) =>
        // Something else happened!
        Promise.reject(err),
      )
  }, [])

  const theme = useTheme()
  const [open, setOpen] = useState(false)

  const handleDrawerOpen = () => {
    setOpen(true)
  }

  const handleDrawerClose = () => {
    setOpen(false)
  }

  return [
    <AppBar key="appbar" position="fixed" open={open}>
      <Toolbar>
        <IconButton
          color="inherit"
          aria-label="menu"
          onClick={handleDrawerOpen}
          edge="start"
          sx={{
            marginRight: 5,
            ...(open && { display: 'none' }),
          }}
          disabled={!hasSession}
        >
          <MenuIcon />
        </IconButton>

        <Link href="/">
          <Button href="/" color="inherit" sx={{ flexGrow: 1, display: { xs: 'none', sm: 'block' } }}>
            <Typography
              component="h1"
              variant="h6"
              color="inherit"
              noWrap
            >
              Shortlink
            </Typography>
          </Button>
        </Link>

        <ToggleDarkMode />

        { secondMenu() }

        <SearchForm />

        {hasSession ? (
          <Fragment>
            <Profile />

            <Notification />
          </Fragment>
        ) : (
          <Link href="/auth/login">
            <Button variant="outlined" color="inherit">
              Log in
            </Button>
          </Link>
        )}
      </Toolbar>
    </AppBar>,
    <Fragment>
      {hasSession && (
        <Drawer key="drawer" variant="permanent" open={open}>
          <DrawerHeader>
            <IconButton onClick={handleDrawerClose}>
              {theme.direction === 'rtl' ? (
                <ChevronRightIcon />
              ) : (
                <ChevronLeftIcon />
              )}
            </IconButton>
          </DrawerHeader>
          <Divider />

          <List>{mainListItems}</List>
          <Divider />

          <List>{secondaryListItems}</List>
          <Divider />

          <List>{adminListItems}</List>
        </Drawer>
      )}
    </Fragment>,
  ]
}

export default Header
