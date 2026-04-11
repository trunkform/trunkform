import { describe, it, expect } from 'vitest'
import { mount } from '@vue/test-utils'
import { readFileSync } from 'fs'
import { fileURLToPath } from 'url'
import { dirname, join } from 'path'
import TrunkformTrademark from './TrunkformTrademark.vue'

const __dirname = dirname(fileURLToPath(import.meta.url))
  
describe('TrunkformTrademark.vue', () => {
  it('renders correctly', () => {
    const wrapper = mount(TrunkformTrademark)
    expect(wrapper.find('p').find('a').text())
      .toBe('trunkform is a trademark of Richard James Minchuk')
    expect(wrapper.find('p').find('a').attributes('href'))
      .toBe('https://tsdr.uspto.gov/#caseNumber=99631818&caseSearchType=US_APPLICATION&caseType=DEFAULT&searchType=statusSearch')
  })

  it('style block has not changed', () => {
    const src = readFileSync(join(__dirname, 'TrunkformTrademark.vue'), 'utf-8')
    const style = src.match(/<style[^>]*>([\s\S]*?)<\/style>/)?.[1]
    expect(style).toMatchSnapshot()
  })
})
