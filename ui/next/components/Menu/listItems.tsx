import React from 'react'
import Link from 'next/link'
import ListItem from '@mui/material/ListItem'
import Tooltip from '@mui/material/Tooltip'
import ListItemIcon from '@mui/material/ListItemIcon'
import ListItemText from '@mui/material/ListItemText'
import ListSubheader from '@mui/material/ListSubheader'
import DashboardIcon from '@mui/icons-material/Dashboard'
import ListIcon from '@mui/icons-material/List'
import BarChartIcon from '@mui/icons-material/BarChart'
import LayersIcon from '@mui/icons-material/Layers'
import AssignmentIcon from '@mui/icons-material/Assignment'
import AccountBalanceWalletIcon from '@mui/icons-material/AccountBalanceWallet'
import AssessmentIcon from '@mui/icons-material/Assessment'
import PersonIcon from '@mui/icons-material/Person'
import HttpIcon from '@mui/icons-material/Http'
import PeopleIcon from '@mui/icons-material/People'
import GroupAddIcon from '@mui/icons-material/GroupAdd'

const mainMenuList = [
  {
    name: 'Add URL',
    url: '/user/addUrl',
    icon: <HttpIcon />,
  },
  {
    name: 'Dashboard',
    url: '/user/dashboard',
    icon: <DashboardIcon />,
  },
  {
    name: 'Links',
    url: '/user/links',
    icon: <ListIcon />,
  },
  {
    name: 'Reports',
    url: '/user/reports',
    icon: <BarChartIcon />,
  },
  {
    name: 'Profile',
    url: '/user/profile',
    icon: <PersonIcon />,
  },
  {
    name: 'Integrations',
    url: '/user/integrations',
    icon: <LayersIcon />,
  },
]

export const mainListItems = mainMenuList.map((item) => (
  <Link href={item.url} key={item.url}>
    <ListItem>
      <Tooltip title={item.name}>
        <ListItemIcon>{item.icon}</ListItemIcon>
      </Tooltip>
      <ListItemText primary={item.name} />
    </ListItem>
  </Link>
))

const otherMenuList = [
  {
    name: 'Billing',
    url: '/user/billing',
    icon: <AccountBalanceWalletIcon />,
  },
  {
    name: 'Audit',
    url: '/user/audit',
    icon: <AssessmentIcon />,
  },
  {
    name: 'About As',
    url: '/about',
    icon: <AssignmentIcon />,
  },
]

export const secondaryListItems = [
  <ListSubheader inset>Other options</ListSubheader>,
  otherMenuList.map((item) => (
    <Link href={item.url} key={item.url}>
      <ListItem>
        <Tooltip title={item.name}>
          <ListItemIcon>{item.icon}</ListItemIcon>
        </Tooltip>
        <ListItemText primary={item.name} />
      </ListItem>
    </Link>
  )),
]

const adminMenuList = [
  {
    name: 'Groups',
    url: '/admin/groups',
    icon: <PeopleIcon />,
  },
  {
    name: 'Users',
    url: '/admin/users',
    icon: <GroupAddIcon />,
  },
  {
    name: 'Links',
    url: '/admin/links',
    icon: <ListIcon />,
  },
]

export const adminListItems = [
  <ListSubheader inset>Admin options</ListSubheader>,
  adminMenuList.map((item) => (
    <Link href={item.url} key={item.url}>
      <ListItem>
        <Tooltip title={item.name}>
          <ListItemIcon>{item.icon}</ListItemIcon>
        </Tooltip>
        <ListItemText primary={item.name} />
      </ListItem>
    </Link>
  )),
]
