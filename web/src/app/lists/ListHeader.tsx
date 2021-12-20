import React, { ReactNode } from 'react'
import { CardHeader, Typography } from '@mui/material'
import makeStyles from '@mui/styles/makeStyles'

const useStyles = makeStyles(() => ({
  headerNote: {
    fontStyle: 'italic',
  },
}))

export function ListHeader(props: {
  cardHeader?: ReactNode
  headerNote?: string
  headerAction?: JSX.Element
}): JSX.Element {
  const classes = useStyles()
  const { headerNote, headerAction, cardHeader } = props
  return (
    <React.Fragment>
      {cardHeader}
      {(headerNote || headerAction) && (
        <CardHeader
          subheader={
            <Typography color='textSecondary' className={classes.headerNote}>
              {headerNote}
            </Typography>
          }
          action={headerAction}
        />
      )}
    </React.Fragment>
  )
}
