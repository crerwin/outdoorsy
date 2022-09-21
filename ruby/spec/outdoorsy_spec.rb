# frozen_string_literal: true

require './outdoorsy'

describe 'determine_delimiter' do
  it 'correctly detects comma delimiters' do
    expect(determine_delimiter('a,b')).to eq(',')
  end
  it 'correctly detects pipe delimiters' do
    expect(determine_delimiter('a|b')).to eq('|')
  end

  # error handling
  it 'correctly raises an error when no delimiters are detected' do
    # TODO: assert error messages
    expect { determine_delimiter('ab') }.to raise_error(Exception)
    expect { determine_delimiter('a.b') }.to raise_error(Exception)
  end
  it 'correctly raises an error when more than one delimiter is detected' do
    expect { determine_delimiter('a,b|c') }.to raise_error(Exception)
  end
end

describe 'load_customer' do
  it 'raises an error when the customer has bad data' do
    expect { load_customer('a,b,c') }.to raise_error(Exception)
  end
end