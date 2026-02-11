import { describe, it, expect } from 'vitest'
import router from './router'

describe('router', () => {
  it('has a Home route at /', () => {
    const route = router.getRoutes().find(r => r.name === 'Home')

    expect(route).toBeTruthy()
    expect(route.path).toBe('/')
  })
})
