import { it, expect } from 'vitest'
import { readFileSync } from 'fs'
import { fileURLToPath } from 'url'
import { dirname, join } from 'path'

const __dirname = dirname(fileURLToPath(import.meta.url))

it('style.css has not changed', () => {
  const css = readFileSync(join(__dirname, 'style.css'), 'utf-8')
  expect(css).toMatchSnapshot()
})
