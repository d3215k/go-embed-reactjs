import { createLazyFileRoute } from '@tanstack/react-router'

import { Label } from "@/components/ui/label"
import { Switch } from "@/components/ui/switch"


export const Route = createLazyFileRoute('/about')({
  component: About,
})

function About() {
  return (
    <div>
      <div className="flex items-center space-x-2">
        <Switch id="airplane-mode" />
        <Label htmlFor="airplane-mode">Airplane Mode</Label>
      </div>
    </div>
    )
}