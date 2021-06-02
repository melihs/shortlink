import React from 'react'
import Link from 'next/link'
import ListItem from '@material-ui/core/ListItem'
import ListItemIcon from '@material-ui/core/ListItemIcon'
import ListItemText from '@material-ui/core/ListItemText'
import ListSubheader from '@material-ui/core/ListSubheader'
import DashboardIcon from '@material-ui/icons/Dashboard'
import ListIcon from '@material-ui/icons/List'
import BarChartIcon from '@material-ui/icons/BarChart'
import LayersIcon from '@material-ui/icons/Layers'
import AssignmentIcon from '@material-ui/icons/Assignment'
import AccountBalanceWalletIcon from '@material-ui/icons/AccountBalanceWallet'
import AssessmentIcon from '@material-ui/icons/Assessment'
import PersonIcon from '@material-ui/icons/Person'

const mainMenuList = [
  {
    name: "Dashboard",
    url: "/user/dashboard",
    icon: <DashboardIcon />,
  },
  {
    name: "Links",
    url: "/user/links",
    icon: <ListIcon />,
  },
  {
    name: "Reports",
    url: "/user/reports",
    icon: <BarChartIcon />,
  },
  {
    name: "Profile",
    url: "/user/profile",
    icon: <PersonIcon />,
  },
  {
    name: "Integrations",
    url: "/user/integrations",
    icon: <LayersIcon />,
  },
]

export const mainListItems = mainMenuList.map(item => (
  <Link href={item.url} key={item.url}>
    <ListItem button>
      <ListItemIcon>
        {item.icon}
      </ListItemIcon>
      <ListItemText primary={item.name} />
    </ListItem>
  </Link>
))

const otherMenuList = [
  {
    name: "Billing",
    url: "/user/billing",
    icon: <AccountBalanceWalletIcon />,
  },
  {
    name: "Audit",
    url: "/user/audit",
    icon: <AssessmentIcon />,
  },
  {
    name: "About As",
    url: "/about",
    icon: <AssignmentIcon />,
  },
]

export const secondaryListItems = [
  <ListSubheader inset>Other options</ListSubheader>,
  otherMenuList.map(item => (
    <Link href={item.url} key={item.url}>
      <ListItem button>
        <ListItemIcon>
          {item.icon}
        </ListItemIcon>
        <ListItemText primary={item.name} />
      </ListItem>
    </Link>
  ))
]
